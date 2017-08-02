package app

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"iosxc.com/levante/util"
	"os"
)

type AppConfig struct {
	Home string

	Server struct {
		Port    uint
		Addr    string
		CharSet string
	}
	Database struct {
		User     string
		Password string
		Schema   string
		Host     string
		Port     uint
	}
	View struct {
		Static struct {
			Path string
			URI  string
		}
		Template struct {
			Layout string
			Path   string
			Ext    string
			Reload bool
		}
	}
	Log struct {
		File        string
		RotateType  string
		RotateValue string
		RotateCount uint
	}
}

var config *AppConfig

func initConfig(application *iris.Application, configPath string) *AppConfig {

	configFile, err := os.Open(configPath)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	application.UseGlobal(func(context context.Context) {
		context.Values().Set(util.CONST_APP_CONFIG, config)
		context.Next()
	})
	return config
}
