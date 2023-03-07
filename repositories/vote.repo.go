package repositories

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stefanusong/votify-api/models"
	"gorm.io/gorm"
)

type VoteRepository interface {
	InsertVote(vote models.Vote) (uuid.UUID, error)
	GetVotesByUserID(userId string) ([]models.Vote, error)
	GetPublicVotes() ([]models.Vote, error)
	GetVoteByID(voteId string) (*models.Vote, error)
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

func (vr *voteRepository) GetVotesByUserID(userId string) ([]models.Vote, error) {
	var votes []models.Vote
	err := vr.db.Select("id", "created_at", "updated_at", "user_id",
		"category_id", "slug", "title", "is_open", "is_public").Find(&votes, "user_id = ?", userId).Error
	return votes, err
}

func (vr *voteRepository) GetPublicVotes() ([]models.Vote, error) {
	// TODO: Apply Pagination
	var votes []models.Vote
	err := vr.db.Select("id", "created_at", "updated_at", "user_id",
		"category_id", "slug", "title", "is_open", "is_public").Find(&votes, "is_public = ?", true).Error
	return votes, err
}

func (vr *voteRepository) GetVoteByID(voteId string) (*models.Vote, error) {
	var vote *models.Vote
	res := vr.db.Select("id", "created_at", "updated_at", "user_id",
		"category_id", "slug", "title", "is_open", "is_public").First(&vote, "ID = ?", voteId)
	err := res.Error

	if res.RowsAffected == 0 {
		return nil, nil
	}

	return vote, err
}
