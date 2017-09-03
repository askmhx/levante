package orm


type LinkGroup struct {
	OprBaseModel
	Title       string
	Description string
	Sort        uint
	Links   	[]Link
}


type Link struct {
	OprBaseModel
	Url         string
	Title       string
	Image       string
	Description string
	Owner       string
	Rating      uint
	LinkGroupID uint
	Visible     bool
	Highlight   bool
	Sort        uint
}