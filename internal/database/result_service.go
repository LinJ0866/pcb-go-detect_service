package database

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func GetAllResults() []results {
	var results []results
	instance.Find(&results)

	for i := 0; i < len(results); i++ {
		timestamp, _ := strconv.ParseInt(results[i].UpdateTime, 10, 64)
		results[i].UpdateTime = time.Unix(timestamp, 0).Format("2006年01月02日 15:04:05")
	}
	return results
}

func AddResult(url string, resultStr string, counts []int) int {
	resultInfo := results{
		PicUrl:     url,
		UpdateTime: fmt.Sprintf("%d", time.Now().Unix()),
		Results:    resultStr,
		Count0:     counts[0],
		Count1:     counts[1],
		Count2:     counts[2],
		Count3:     counts[3],
		Count4:     counts[4],
	}
	err := instance.Create(&resultInfo).Error
	if err != nil {
		log.Println("create new resultInfo err: ", err)
		return -1
	}
	return 0
}
