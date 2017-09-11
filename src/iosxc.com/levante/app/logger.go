package app

import (
	"time"
	"os"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/context"
	"fmt"
)


func newLogFile(config *AppConfig) *os.File {
	logPath := fmt.Sprintf("%s%s", config.Home, config.Log.File)
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

func NewRequestLogger(config *AppConfig) (h context.Handler) {
	c := logger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Columns: true,
	}
	logFile := newLogFile(config)
	c.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}) {
		output := logger.Columnize(now.Format("2006/01/02 - 15:04:05"), latency, status, ip, method, path, message)
		logFile.Write([]byte(output))
	}
	h = logger.New(c)
	return
}
