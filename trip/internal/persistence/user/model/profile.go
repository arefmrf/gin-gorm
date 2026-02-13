package model

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type NotificationLanguage string

const (
	LangPersian NotificationLanguage = "fa"
	LangEnglish NotificationLanguage = "en"
	LangArabic  NotificationLanguage = "ar"
)

type Profile struct {
	ID uint `gorm:"primaryKey"`

	UserID uint `gorm:"uniqueIndex;not null"` // OneToOne
	User   User

	Avatar    *string `gorm:"size:255"`
	FirstName *string `gorm:"size:32"`
	LastName  *string `gorm:"size:32"`
	Bio       *string `gorm:"type:text"`

	NotificationLanguage NotificationLanguage `gorm:"size:5;default:'fa'"`

	Extra datatypes.JSON `gorm:"type:jsonb"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (p *Profile) TableName() string {
	return "profiles"
}

func defaultProfileExtra() datatypes.JSON {
	data := map[string]any{
		"chat":            map[string]any{},
		"integration":     nil,
		"has_unread_chat": false,
		"social_network":  map[string]any{},
		"iso2":            nil,
		"bank_info":       map[string]any{},
	}

	bytes, _ := json.Marshal(data)
	return datatypes.JSON(bytes)
}

func (p *Profile) AvatarURL() string {
	return GetProxyURL(p.Avatar)
}

func (p *Profile) BeforeCreate(tx *gorm.DB) error {
	if len(p.Extra) == 0 {
		p.Extra = defaultProfileExtra()
	}
	return nil
}

func GenerateAvatarPath(filename string) string {
	ext := filepath.Ext(filename)
	id := ulid.Make().String()
	return fmt.Sprintf("avatars/%s%s", id, ext)
}

//path := GenerateAvatarPath(fileHeader.Filename)
//profile.Avatar = &path

func GetProxyURL(objectName *string) string {
	if objectName == nil || *objectName == "" {
		return ""
	}

	proxyURL := os.Getenv("S3_ENDPOINT_PROXY_URL")
	if proxyURL == "" {
		panic("S3_ENDPOINT_PROXY_URL not set")
	}

	if !strings.HasPrefix(proxyURL, "http") {
		scheme := "https"
		if os.Getenv("DEBUG") == "1" {
			scheme = "http"
		}
		proxyURL = scheme + "://" + proxyURL
	}

	if !strings.HasSuffix(proxyURL, "/") {
		proxyURL += "/"
	}

	return proxyURL + *objectName
}
