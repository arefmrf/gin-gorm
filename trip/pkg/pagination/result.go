package pagination

import "math"

type Result[T any] struct {
	Data       T     `json:"data"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

func NewResult[T any](data T, p Pagination, total int64) Result[T] {
	return Result[T]{
		Data:       data,
		Page:       p.Page,
		Limit:      p.Limit,
		Total:      total,
		TotalPages: int(math.Ceil(float64(total) / float64(p.Limit))),
	}
}
