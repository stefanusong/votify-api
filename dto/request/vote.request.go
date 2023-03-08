package request

import uuid "github.com/satori/go.uuid"

type CreateVote struct {
	UserID     uuid.UUID        // not required
	CategoryID uuid.UUID        `json:"category_id" binding:"required"`
	Slug       string           `json:"slug" binding:"required,min=3"`
	Title      string           `json:"title" binding:"required,min=3"`
	IsOpen     bool             `json:"is_open"`
	IsPublic   bool             `json:"is_public"`
	Questions  []CreateQuestion `json:"questions" binding:"required"`
}

type UserAnswer struct {
	VoteID  uuid.UUID
	UserID  uuid.UUID
	Answers []Answer
}

type Answer struct {
	QuestionID uuid.UUID
	OptionID   uuid.UUID
}
type CreateQuestion struct {
	Question string         `json:"question" binding:"required,min=3"`
	Options  []CreateOption `json:"options" binding:"required"`
}

type CreateOption struct {
	Option string `json:"option" binding:"required,min=3"`
	Desc   string `json:"desc"`
}
