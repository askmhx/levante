package orm

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	OprBaseModel
	Email         string
	UserName      string
	Avatar        string
	Mobile        string
	Site          string
	Bio           string
	Birthday      string
	Password      string
	IsStaff       bool
	IsActive      bool
	IsAdmin       bool
	LoginCount    uint64
	LastLoginIp   string
	LastLoginAt   time.Time
	ErrLoginCount uint
}
