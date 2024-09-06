package token

import (
	"encoding/json"
	"fmt"
	"go-detect_service/config"
	"io"
	"log"
	"net/http"
	"time"
)

func GetNewToken() (int, string) {
	var client = &http.Client{
		Timeout: time.Second * 5,
	}
	rqst, err := client.Get(fmt.Sprintf("%s/huawei/getToken", config.Cfg.ApiURL))
	if err != nil {
		log.Println("New request failed:", err)
		return -1, ""
	}

	resText, _ := io.ReadAll(rqst.Body)
	var tokenRes tokenResStruct
	err = json.Unmarshal(resText, &tokenRes)
	if err != nil {
		log.Println("parse resBody error: ", err)
	}

	return 0, tokenRes.Data.Token
}
