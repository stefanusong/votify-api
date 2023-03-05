package models

import uuid "github.com/satori/go.uuid"

type VoteQuestion struct {
	Base
	VoteID   uuid.UUID `gorm:"not null"`
	Question string    `gorm:"not null"`

	VoteOptions []VoteOption `gorm:"foreignKey:VoteQuestionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func NewVoteQuestion(question string, voteOptions []VoteOption) VoteQuestion {
	return VoteQuestion{
		Question:    question,
		VoteOptions: voteOptions,
	}
}
