package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"go-detect_service/config"
	"go-detect_service/internal/controller/detect"
	"go-detect_service/internal/controller/qiniu"
	"go-detect_service/internal/controller/token"
	"go-detect_service/internal/database"
	"go-detect_service/internal/utils"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func testResponse(c *gin.Context) {
	c.JSON(504, "timeout")
}

func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(30*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(testResponse),
	)
}

func InitRouter() {
	gin.SetMode(config.Cfg.AppMode)
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(timeoutMiddleware())

	TokenRouter := router.Group("/detect")
	{
		TokenRouter.GET("/", func(c *gin.Context) {
			c.JSON(200, utils.SendResult(200, fmt.Sprintf("server@%s start successful!", config.Version), nil))
		})

		TokenRouter.POST("/infer", inferRouter)
		TokenRouter.GET("/history", func(c *gin.Context) {
			results := database.GetAllResults()
			c.JSON(200, utils.SendResult(200, "获取成功", results))
		})
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(200, utils.SendResult(404, fmt.Sprintf("go-detect_service@%s server: not found", config.Version), nil))
	})

	router.Run(config.Cfg.Port)
}

func inferRouter(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Println("gin get file error: ", err)
		c.JSON(400, "检查参数")
		return
	}
	log.Println(header.Filename)

	// 获取token
	code, token := token.GetNewToken()
	if code != 0 {
		c.JSON(400, "服务获取token获取失败")
		return
	}

	// 设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Cfg.Timeout))
	defer cancel()

	var id_ int
	id := make(chan int)
	err_endless := make(chan error)
	go detect.GetAvailableService(ctx, id, err_endless)
	select {
	case res := <-id:
		id_ = res
		break
	case <-err_endless:
		c.JSON(504, "获取空余service超时")
		return
	}

	serviceInfo := database.GetServiceInfo(id_)

	// return
	code, resData := detect.Detect(serviceInfo.Url, token, file)
	if code != 0 {
		c.JSON(4003, "推理服务失败")
		return
	}
	m1 := make(map[string][]interface{})
	_ = json.Unmarshal(resData, &m1)

	c.JSON(200, m1)

	database.UpdateStatus(id_, 0)

	code, url := qiniu.UploadToQiNiu(header)
	if code != 0 {
		return
	}

	detect.CreateNewResult(url, m1["detection_classes"], string(resData))
}
