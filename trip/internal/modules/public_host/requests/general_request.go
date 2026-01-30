package requests

type PublicHostListRequest struct {
	Search         *string  `json:"search"`
	StartDate      *int64   `json:"start_date"`
	EndDate        *int64   `json:"end_date"`
	Location       *string  `json:"location"`
	MaleCapacity   *int     `json:"male_capacity"`
	FemaleCapacity *int     `json:"female_capacity"`
	Capacity       *int     `json:"capacity"`
	LowestPrice    *float64 `json:"lowest_price"`
	HighestPrice   *float64 `json:"highest_price"`
	HostType       *string  `json:"host_type"`
	Category       *string  `json:"category"`
	Facilities     *string  `json:"facilities"`
	HasImage       *bool    `json:"has_image"`
	Order          *string  `json:"order"`
	Tags           *string  `json:"tags"`
}
