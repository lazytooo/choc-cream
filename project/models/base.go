package models

type PageInfo struct {
	Page     int64
	PageSize int64
	UsePage  bool
}

func (p *PageInfo) Limit() (start, size int64) {
	if !p.UsePage {
		return
	}
	start = (p.Page - 1) * p.PageSize

	return start, p.PageSize
}
