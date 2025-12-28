package model

import "trip/internal/persistence"

type Status int16
type CancelLevel int16

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

type Host struct {
	persistence.BaseModel

	PlaceID *uint
	Place   *Place `gorm:"foreignKey:PlaceID;references:ID"`

	HostTypeID *uint     `gorm:"index"`
	HostType   *HostType `gorm:"foreignKey:HostTypeID;references:ID"`

	UID  string `gorm:"size:26;uniqueIndex"`
	User string `gorm:"size:26;index"`

	Title                    string `gorm:"size:128"`
	Description              string
	Hidden                   bool
	BookingPaymentExpireTime *int16
	Priority                 int16
	Online                   bool
	OutOfService             bool
	Info                     persistence.JSONB `gorm:"type:jsonb"`
	Status                   Status            `gorm:"type:smallint;default:1"`

	Facilities []persistence.Facility `gorm:"many2many:host_facilities"`
	Tags       []persistence.HostTag  `gorm:"many2many:host_tags"`

	CreatedAt  int64  `gorm:"autoCreateTime"`
	UpdatedAt  int64  `gorm:"autoUpdateTime"`
	Identifier string `gorm:"size:16;uniqueIndex"`
	Rate       float64
	RateCount  int16
	//CurrencyID  *uint
	CancelLevel CancelLevel `gorm:"type:smallint;default:1"`
}

func (Host) TableName() string { return "hosts" }
