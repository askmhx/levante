package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"iosxc.com/levante/orm"
	"iosxc.com/levante/util"
	"log"
	"os"
)

var db *gorm.DB

func initDatabase(config *AppConfig) *gorm.DB{
	var err error
	dbConnURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Schema)
	db, err = gorm.Open("mysql", dbConnURL)
	if err != nil {
		panic("failed to connect database")
	}
	setDBLogger(config)
	defer db.Close()
	db.AutoMigrate(&orm.Post{})
	db.AutoMigrate(&orm.Link{})
	db.AutoMigrate(&orm.User{})
	return db
}

func setDBLogger(config *AppConfig) {
	logPath := fmt.Sprintf("%s%s", config.Home, config.Log.File)
	if !util.CheckIsExistPath(logPath) {
		panic("logPath:" + logPath + " is not exist! bootstrap.go must execute setLogger() before setDatabase()")
	}
	file, _ := os.Open(logPath)
	db.SetLogger(log.New(file, "\r\nDB:", 0))
}
