package app

import (
	"time"
	"os"
	"github.com/kataras/iris/middleware/logger"
	"strings"
	"github.com/kataras/iris/context"
	"fmt"
)

const deleteFileOnExit = false

var excludeExtensions = [...]string{
	".js",
	".css",
	".jpg",
	".png",
	".ico",
	".svg",
}

func newLogFile(config *AppConfig) *os.File {
	logPath := fmt.Sprintf("%s%s", config.Home, config.Log.File)
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

func NewRequestLogger(config *AppConfig) (h context.Handler, close func() error) {
	close = func() error { return nil }
	c := logger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Columns: true,
	}
	logFile := newLogFile(config)
	close = func() error {
		err := logFile.Close()
		if deleteFileOnExit {
			err = os.Remove(logFile.Name())
		}
		return err
	}

	c.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}) {
		output := logger.Columnize(now.Format("2006/01/02 - 15:04:05"), latency, status, ip, method, path, message)
		logFile.Write([]byte(output))
	}

	c.AddSkipper(func(ctx context.Context) bool {
		path := ctx.Path()
		for _, ext := range excludeExtensions {
			if strings.HasSuffix(path, ext) {
				return true
			}
		}
		return false
	})

	h = logger.New(c)

	return
}
