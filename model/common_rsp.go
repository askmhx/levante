package model

type SiteUrl struct {
	Loc      string
	LastMod  string
	Priority int
}

type ColCount struct {
	ColName string
	Count   int
}

type RspData struct {
	RspCode string
	RspMsg  string
	RspURL  string
}

type LinkData struct {
	Id          int
	Title       string
	Url         string
	Image       string
	Description string
	Owner       string
	Highlight   int
	Gtitle      string
}

type LinkGroupEntry struct {
	Title    string
	LinkList []LinkData
}

var RspDataUnknowError = RspData{RspCode: "900500", RspMsg: "未知错误", RspURL: "/"}

var RspDataPageNotFound = RspData{RspCode: "900404", RspMsg: "找不到访问链接", RspURL: "/"}
