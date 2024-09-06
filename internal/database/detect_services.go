package database

import (
	"log"
)

func GetServiceInfo(id int) detectServices {
	var serviceInfo detectServices
	instance.Select("id", "url").Where("id = ?", id).First(&serviceInfo)
	return serviceInfo
}

func GetAvailableService() (int, detectServices) {
	var serviceInfo detectServices
	result := instance.Select("id", "url").Where("status = 0 and is_online = 0").First(&serviceInfo)
	if result.Error != nil {
		log.Println("err:", result.Error)
		return -1, serviceInfo
	}
	return 0, serviceInfo
}

func UpdateStatus(id int, status int) int {
	statusOption := ""
	if status == -1 {
		statusOption = " and status = 0"
	}

	err := instance.Model(&detectServices{}).Where("id = ?"+statusOption, id).Update("status", status).Error
	if err != nil {
		log.Println("update detectService status err: ", err)
		return -1
	}

	return 0
}
