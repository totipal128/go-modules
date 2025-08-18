package postgreSQL

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
