package model

type Line struct {
	ID    uint `gorm:"primaryKey"`
	Valid bool
}
