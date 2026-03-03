package model

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type PaginatedResponse[T any] struct {
	Data       []T        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func (p *Pagination) Check() {
	if p.Page < 1 {
		p.Page = 1
	}

	if p.Limit < 1 || p.Limit > 50 {
		p.Limit = 10
	}
}

func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}
