package app

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"os"
)

type AppConfig struct {
	Home    string
	IrisYML string

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
		Level       string
	}
}

var config *AppConfig

func InitConfig(app *iris.Application, configPath string) *AppConfig {
	configFile, err := os.Open(configPath)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	app.Configure(iris.WithConfiguration(iris.YAML(config.IrisYML)))
	return config
}
