package orm

type Post struct {
	OprBaseModel
	Title         string `gorm:"primary_key"`
	Content       string
	Author        string
	Support       uint64
	Oppose        uint64
	Views         uint64
	PostStatus    uint64
	CommentStatus uint64
	PermanentURL  string
	PostType      string
	SourceLink    string
	Catalog string
}
