package response

var (
	pageIndexMin = 1

	pageSizeMin = 1
	pageSizeDef = 10
	pageSizeMax = 500

	Page = &page{}
)

type PageInfo struct {
	Pagination
	Total   int64       `json:"total"`
	Records interface{} `json:"records"`
}

type Pagination struct {
	PageIndex int `form:"pageIndex" json:"pageIndex"`
	PageSize  int `form:"pageSize" json:"pageSize"`
}

type page struct {
}

func NewPage(total int64, pagination *Pagination, records interface{}) *PageInfo {
	return &PageInfo{Total: total, Pagination: *Page.GetPagination(pagination), Records: records}
}

func NilPage(pagination *Pagination) *PageInfo {
	return NewPage(0, pagination, make([]struct{}, 0))
}

func (p *page) GetPage(pagination *Pagination) (int, int) {
	page := p.GetPagination(pagination)
	return page.PageIndex, page.PageSize
}

func (p *page) GetPagination(pagination *Pagination) *Pagination {

	if pagination == nil {
		return &Pagination{PageIndex: pageIndexMin, PageSize: pageSizeDef}
	}

	pageIndex := pagination.PageIndex
	if pagination.PageIndex < pageIndexMin {
		pageIndex = pageSizeMin
	}

	pageSize := pagination.PageSize
	if pageSize < pageSizeMin {
		pageSize = pageSizeMin
	} else if pageSize > pageSizeMax {
		pageSize = pageSizeMax
	}

	return &Pagination{PageIndex: pageIndex, PageSize: pageSize}
}

func (p *page) GetOffset(pagination *Pagination) (int, int) {
	pageIndex, pageSize := p.GetPage(pagination)
	offset := (pageIndex - 1) * pageSize
	return offset, pageSize
}
