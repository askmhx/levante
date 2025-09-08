package model

type Page struct {
	PageIndex int
	PageSize  int
}

func (this Page) Start() int {
	return this.PageIndex * this.PageSize
}

func (this Page) End() int {
	return (this.PageIndex + 1) * this.PageSize
}
