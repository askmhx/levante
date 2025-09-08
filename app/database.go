package app

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"levante/orm"
	"levante/util"
	"log"
	"os"
)

var db *gorm.DB

func InitDatabase(config *AppConfig) *gorm.DB {
	var err error
	dbConnURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Schema)
	db, err = gorm.Open(mysql.Open(dbConnURL), &gorm.Config{Logger: getLogger(config)})
	if err != nil {
		panic("failed to connect database")
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&orm.Post{})
	db.AutoMigrate(&orm.LinkGroup{})
	db.AutoMigrate(&orm.Link{})
	db.AutoMigrate(&orm.User{})
	return db
}

func getLogger(config *AppConfig) logger.Interface {
	logPath := fmt.Sprintf("%s%s", config.Home, config.ApplicationLog.File)
	if !util.CheckIsExistPath(logPath) {
		panic("logPath:" + logPath + " is not exist! bootstrap.go must execute setLogger() before setDatabase()")
	}
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return logger.New(
		log.New(logFile, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
}
