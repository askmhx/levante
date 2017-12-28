package model

type Page struct {
	PageIndex uint64
	PageSize  uint64
}

func (this Page) Start() uint64 {
	return this.PageIndex * this.PageSize
}

func (this Page) End() uint64 {
	return (this.PageIndex + 1) * this.PageSize
}
