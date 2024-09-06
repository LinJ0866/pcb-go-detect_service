package detect

import (
	"bytes"
	"context"
	"fmt"
	"go-detect_service/internal/database"
	"go-detect_service/internal/utils"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"
)

func CreateNewResult(pic_url string, detectionClasses []interface{}, resText string) {
	counts := []int{0, 0, 0, 0, 0}
	classList := []string{
		"Mouse_bite",
		"Open_circuit",
		"Short",
		"Spur",
		"Spurious_copper",
	}

	for _, resultItem := range detectionClasses {
		str := resultItem.(string)
		index := utils.Find(str, classList)
		if index == -1 {
			log.Println("error className: ", str)
			continue
		}
		counts[index] += 1
	}

	database.AddResult(pic_url, resText, counts)
}

func GetAvailableService(ctx context.Context, id chan<- int, err chan<- error) {
	for {
		select {
		case <-ctx.Done():
			err <- fmt.Errorf("operation timed out")
			return
		default:
			code, serviceInfo := database.GetAvailableService()
			if code == 0 {
				if code1 := database.UpdateStatus(serviceInfo.Id, -1); code1 == 0 {
					id <- serviceInfo.Id
					return
				}
			}
		}
	}
}

func Detect(url string, token string, file multipart.File) (int, []byte) {
	var client = &http.Client{
		Timeout: time.Second * 30,
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", fmt.Sprintf("%d.jpg", time.Now().Unix()))
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		log.Println("copy content to struct multipart error: ", err)
		return -1, []byte("")
	}
	writer.Close()

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println("create httpRequest failed:", err)
		return -1, []byte("")
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Auth-Token", token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("new Request failed:", err)
		return -1, []byte("")
	}
	resText, _ := io.ReadAll(resp.Body)

	return 0, resText
}
