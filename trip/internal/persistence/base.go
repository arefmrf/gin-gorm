package persistence

type BaseModel struct {
	ID uint `gorm:"primaryKey"`
}
