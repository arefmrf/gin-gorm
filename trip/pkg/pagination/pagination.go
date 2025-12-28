package pagination

type Pagination struct {
	Page  int
	Limit int
}

const (
	DefaultPage  = 1
	DefaultLimit = 100
	MaxLimit     = 500
)

func New(page, limit int) Pagination {
	if page <= 0 {
		page = DefaultPage
	}
	if limit <= 0 {
		limit = DefaultLimit
	}
	if limit > MaxLimit {
		limit = MaxLimit
	}

	return Pagination{
		Page:  page,
		Limit: limit,
	}
}

func (p Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}
