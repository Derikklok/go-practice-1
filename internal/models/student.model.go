package models

// Models
type Student struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Grade int    `gorm:"default:0"`
}
