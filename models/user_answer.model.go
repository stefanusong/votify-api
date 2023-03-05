package models

import uuid "github.com/satori/go.uuid"

type UserAnswer struct {
	Base
	VoteID     uuid.UUID `gorm:"not null"`
	UserID     uuid.UUID `gorm:"not null"`
	QuestionID uuid.UUID `gorm:"not null"`
	OptionID   uuid.UUID `gorm:"not null"`

	Vote         Vote         `gorm:"foreignKey:VoteID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User         User         `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	VoteQuestion VoteQuestion `gorm:"foreignKey:QuestionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	VoteOption   VoteOption   `gorm:"foreignKey:OptionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
