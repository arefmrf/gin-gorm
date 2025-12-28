package persistence

type HostTag struct {
	BaseModel
	UID         string `gorm:"size:26;uniqueIndex"`
	Title       string `gorm:"size:255"`
	Translation JSONB  `gorm:"type:jsonb"`
	Status      int16
}

func (HostTag) TableName() string { return "host_tags" }
