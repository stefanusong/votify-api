package models

type Category struct {
	Base
	Name string `gorm:"not null"`

	Votes []Vote `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
