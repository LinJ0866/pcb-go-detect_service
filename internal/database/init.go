package database

import (
	"fmt"
	"go-detect_service/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
)

func InitMysql() {
	conf := config.Cfg.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Pass,
		conf.Address,
		conf.Port,
		conf.DbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	if db, err := db.DB(); err != nil {
		log.Panic(err)
	} else {
		err = db.Ping()
		if err != nil {
			log.Panic(err)
		}
		log.Println("connect to MySQL success")
	}
	instance = db.Debug()

	_ = instance.AutoMigrate(detectServices{})
	_ = instance.AutoMigrate(results{})
}
