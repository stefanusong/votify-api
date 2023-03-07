package response

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stefanusong/votify-api/models"
)

type BaseVote struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserID     uuid.UUID
	CategoryID uuid.UUID
	Slug       string
	Title      string
	IsOpen     bool
	IsPublic   bool
}

func NewBaseVote(vote models.Vote) BaseVote {
	return BaseVote{
		ID:         vote.ID,
		CreatedAt:  vote.CreatedAt,
		UpdatedAt:  vote.UpdatedAt,
		UserID:     vote.UserID,
		CategoryID: vote.CategoryID,
		Slug:       vote.Slug,
		Title:      vote.Title,
		IsOpen:     vote.IsOpen,
		IsPublic:   vote.IsPublic,
	}
}
