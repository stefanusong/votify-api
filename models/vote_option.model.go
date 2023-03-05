package models

import uuid "github.com/satori/go.uuid"

type VoteOption struct {
	Base
	VoteQuestionID uuid.UUID `gorm:"not null"`
	OptionColorID  uuid.UUID
	Option         string `gorm:"not null"`
	Desc           string
	Image          string

	OptionColor OptionColor `gorm:"foreignKey:OptionColorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewVoteOption(optionColorId uuid.UUID, option, desc, image string) VoteOption {
	return VoteOption{
		OptionColorID: optionColorId,
		Option:        option,
		Desc:          desc,
		Image:         image,
	}
}
