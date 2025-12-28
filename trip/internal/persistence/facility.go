package persistence

type Facility struct {
	BaseModel
	UID         string `gorm:"size:26;uniqueIndex"`
	Title       string `gorm:"size:128"`
	Translation JSONB  `gorm:"type:jsonb"`
	Description *string
	Icon        *string
	Active      bool
	Priority    int16
}

func (Facility) TableName() string { return "facilities" }
