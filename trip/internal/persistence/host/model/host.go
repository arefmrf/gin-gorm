package model

import "trip/internal/persistence"

type Status int16
type CancelLevel int16
type PaymentMethod int16

const (
	StatusDraft      Status = 1
	StatusPending    Status = 2
	StatusRejected   Status = 3
	StatusApproved   Status = 4
	StatusSuspension Status = 5
	StatusDeleted    Status = 6
)

const (
	CancelEasy   CancelLevel = 1
	CancelNormal CancelLevel = 2
	CancelHard   CancelLevel = 3
)

const (
	PaymentNone   PaymentMethod = 1
	PaymentCash   PaymentMethod = 2
	PaymentOnline PaymentMethod = 3
	PaymentBoth   PaymentMethod = 4
)

type Host struct {
	persistence.BaseModel

	// Relations
	PlaceID *uint
	Place   *Place `gorm:"foreignKey:PlaceID;references:ID"`

	HostTypeID *uint
	HostType   *HostType `gorm:"foreignKey:HostTypeID;references:ID"`

	//CurrencyID *uint
	//Currency   *Currency `gorm:"foreignKey:CurrencyID;references:ID"`

	// Identity
	UID        string `gorm:"size:26;uniqueIndex"`
	User       string `gorm:"size:26;index"`
	Identifier string `gorm:"size:16;uniqueIndex"`

	// Content
	Title       string `gorm:"size:128"`
	Description string

	// Flags
	Hidden       bool
	Online       bool
	OutOfService bool

	// Booking
	BookingPaymentExpireTime *int16
	Priority                 int16
	Status                   Status        `gorm:"type:smallint;default:1"`
	CancelLevel              CancelLevel   `gorm:"type:smallint;default:1"`
	PaymentMethod            PaymentMethod `gorm:"type:smallint;default:1"`

	// Ratings
	Rate      float64
	RateCount int16

	// Financial
	Subsidy int16

	// Meta / JSON
	Info persistence.JSONB `gorm:"type:jsonb"`

	// M2M
	Facilities []persistence.Facility `gorm:"many2many:host_facilities"`
	Tags       []persistence.HostTag  `gorm:"many2many:host_tags"`

	// Timestamps (unix)
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}

func (Host) TableName() string { return "hosts" }
