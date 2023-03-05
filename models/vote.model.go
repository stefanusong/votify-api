package models

import uuid "github.com/satori/go.uuid"

type Vote struct {
	Base
	UserID     uuid.UUID `gorm:"not null"`
	CategoryID uuid.UUID
	Slug       string `gorm:"not null;index;unique"`
	Title      string `gorm:"not null"`
	IsOpen     bool   `gorm:"not null;default:false"`
	IsPublic   bool   `gorm:"not null;default:false"`
	UseOTP     bool   `gorm:"not null;default:false"`

	User          User           `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Category      Category       `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VoteQuestions []VoteQuestion `gorm:"foreignKey:VoteID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func NewVote(userId, categoryId uuid.UUID, slug, title string, isOpen, isPublic, useOTP bool, voteQuestions []VoteQuestion) Vote {
	return Vote{
		UserID:        userId,
		CategoryID:    categoryId,
		Slug:          slug,
		Title:         title,
		IsOpen:        isOpen,
		IsPublic:      isPublic,
		UseOTP:        useOTP,
		VoteQuestions: voteQuestions,
	}
}
