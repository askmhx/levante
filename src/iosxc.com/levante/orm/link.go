package orm


type LinkGroup struct {
	OprBaseModel
	Title       string
	Description string
	Sort        uint
}


type Link struct {
	OprBaseModel
	Url         string
	Title       string
	Image       string
	Description string
	Owner       string
	Rating      uint
	LinkGroup 	LinkGroup
	LinkGroupID uint
	Visible     bool
	Highlight   bool
	Sort        uint
}