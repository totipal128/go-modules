package mongoDBV2

type Pagination struct {
	//Search
	Keyword   string `json:"-" form:"keyword"`
	StartDate string `json:"-" form:"start_date"`
	EndDate   string `json:"-" form:"end_date"`
	Sort      string `json:"-" form:"sort"`
	SortBy    string `json:"-" form:"sort_by"`

	//Pagination
	ShowAll   bool  `json:"-" form:"show_all"`
	Page      int64 `json:"current_page" form:"page"`
	PageSize  int64 `json:"page_size" form:"page_size"`
	TotalPage int64 `json:"page_total"`
	TotalData int64 `json:"data_total"`
	HasNext   bool  `json:"has_next"`
	HasPrev   bool  `json:"has_previous"`
}

func (p *Pagination) collectDefault() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	if p.SortBy == "" {
		p.SortBy = "id"
	}
	if p.Sort == "" {
		p.Sort = "DESC"
	}
}
