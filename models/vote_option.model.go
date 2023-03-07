package models

import uuid "github.com/satori/go.uuid"

type VoteOption struct {
	Base
	VoteQuestionID uuid.UUID `gorm:"not null"`
	Option         string    `gorm:"not null"`
	Desc           string
}

func NewVoteOption(option, desc string) VoteOption {
	return VoteOption{
		Option: option,
		Desc:   desc,
	}
}
