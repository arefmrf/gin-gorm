package model

import "trip/internal/persistence"

type HostDaily struct {
	persistence.BaseModel

	UID    string `gorm:"size:26;uniqueIndex"`
	HostID uint   `gorm:"index"`

	Date int64 `gorm:"type:int;index"`

	Price      float64
	ExtraPrice float64
	FinalPrice float64

	Discount int16 `gorm:"default:0"`

	Capacity      persistence.JSONB `gorm:"type:jsonb"`
	TotalCapacity int16

	MaleBooked   int16 `gorm:"default:0"`
	FemaleBooked int16 `gorm:"default:0"`

	CreatedAt int64 `gorm:"type:int"`
	UpdatedAt int64 `gorm:"type:int"`
}

func (HostDaily) TableName() string {
	return "host_daily"
}
