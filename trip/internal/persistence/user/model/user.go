package model

import (
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type UserStatus int

const (
	UserStatusActive  UserStatus = 1
	UserStatusBlocked UserStatus = 2
	UserStatusDeleted UserStatus = 3
)

type User struct {
	ID uint `gorm:"primaryKey"`

	UID         string  `gorm:"size:26;uniqueIndex;not null"`
	Username    *string `gorm:"size:26"`
	Password    *string `gorm:"size:128"`
	PhoneNumber *string `gorm:"size:18"`
	Email       *string `gorm:"size:255"`

	Status UserStatus `gorm:"default:1"`

	IsNGO  bool `gorm:"default:false"`
	IsHost bool `gorm:"default:false"`

	LastLogin *time.Time
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Profile *Profile `gorm:"constraint:OnDelete:CASCADE;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.UID == "" {
		u.UID = ulid.Make().String()
	}
	return nil
}

func (u *User) TableName() string {
	return "users"
}
