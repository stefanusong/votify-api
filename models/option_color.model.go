package models

type OptionColor struct {
	Base
	ColorCode string `gorm:"not null"`
	ColorName string `gorm:"not null"`
}
