package Common

type Paging struct {
	Page  int   `json:"page" from:"page"`
	Limit int   `json:"limit" from:"limit"`
	Total int64 `json:"total" from:"-"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 || p.Limit >= 100 {
		p.Limit = 10
	}

}
