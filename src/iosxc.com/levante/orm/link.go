package orm

type Link struct {
	OprBaseModel
	Url         string
	Title       string
	Image       string
	Description string
	Owner       string
	Rating      uint
	LinkGroup   LinkGroup
	Visible     bool
	Highlight   bool
	Sort        uint
}


type LinkGroup struct {
	OprBaseModel
	Title       string
	Description string
	Sort        uint
}
