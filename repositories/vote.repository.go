package repositories

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stefanusong/votify-api/models"
	"gorm.io/gorm"
)

type VoteRepository interface {
	InsertVote(vote models.Vote) (uuid.UUID, error)
}

type voteRepository struct {
	db *gorm.DB
}

func NewVoteRepository(db *gorm.DB) VoteRepository {
	return &voteRepository{
		db: db,
	}
}

func (vr *voteRepository) InsertVote(vote models.Vote) (uuid.UUID, error) {
	err := vr.db.Create(&vote).Error
	return vote.ID, err
}
