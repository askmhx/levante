package orm

import (
	"time"
)

type User struct {
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
