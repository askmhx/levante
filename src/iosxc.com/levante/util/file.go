package util

import (
	"os"
)

func CheckIsExistPath(path string) bool {

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}
