package data

type Pagination struct {
	Limit      int         `json:"limit"`
	Page       int         `json:"page"`
	Sort       string      `json:"sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	SearchKey  string      `json:"search-key"`
	SeachValue string      `json:"search-value"`
	Rows       interface{} `json:"rows"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}
func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "created_at desc"
	}
	return p.Sort
}

func (p *Pagination) GetSearch() string {
	if p.SeachValue != "" {
		if p.SearchKey == "categories" {
			return "categories.name LIKE '%" + p.SeachValue + "%'"
		}

		return p.SearchKey + " LIKE '%" + p.SeachValue + "%'"
	}

	return ""
}
