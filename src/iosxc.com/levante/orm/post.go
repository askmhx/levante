package orm

type Post struct {
	OprBaseModel
	Title      string `gorm:"primary_key"`
	Content    string
	Author     string
	Support    uint64
	Oppose     uint64
	SourceLink string
}
