package repositories

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stefanusong/votify-api/dto/response"
	"github.com/stefanusong/votify-api/models"
	"gorm.io/gorm"
)

type VoteRepository interface {
	InsertVote(vote models.Vote) (uuid.UUID, error)
	GetVotesByUserID(userId string) ([]response.BaseVote, error)
	GetPublicVotes() ([]response.BaseVote, error)
	GetVoteByID(voteId string) (*response.BaseVote, error)
	GetVoteQuestions(voteId string) ([]models.VoteQuestion, error)
	GetQuestionOptions(questionId string) ([]models.VoteOption, error)
	InsertAnswers(answers []models.UserAnswer) error
	GetUserAnswers(voteId, userId string) ([]response.QuestionAnswer, error)
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

func (vr *voteRepository) GetVotesByUserID(userId string) ([]response.BaseVote, error) {
	var votes []response.BaseVote
	err := vr.db.Table("votes").Select("id", "created_at", "updated_at", "user_id",
		"category_id", "slug", "title", "is_open", "is_public").Where("user_id = ?", userId).Scan(&votes).Error
	return votes, err
}

func (vr *voteRepository) GetPublicVotes() ([]response.BaseVote, error) {
	var votes []response.BaseVote
	err := vr.db.Table("votes").Select("id", "created_at", "updated_at", "user_id",
		"category_id", "slug", "title", "is_open", "is_public").Where("is_public = ?", true).Scan(&votes).Error
	return votes, err
}

func (vr *voteRepository) GetVoteByID(voteId string) (*response.BaseVote, error) {
	var vote *response.BaseVote
	res := vr.db.Table("votes").Select("id", "created_at", "updated_at", "user_id",
		"category_id", "slug", "title", "is_open", "is_public").Where("ID = ?", voteId).Scan(&vote)

	if res.RowsAffected == 0 {
		return nil, nil
	}

	return vote, res.Error
}

func (vr *voteRepository) GetVoteQuestions(voteId string) ([]models.VoteQuestion, error) {
	var questions []models.VoteQuestion

	err := vr.db.Table("vote_questions").Select("id", "question").Find(&questions, "vote_id = ?", voteId).Error

	return questions, err
}

func (vr *voteRepository) GetQuestionOptions(questionId string) ([]models.VoteOption, error) {
	var options []models.VoteOption

	err := vr.db.Table("vote_options").Select("id", "option", "desc").Find(&options, "vote_question_id = ?", questionId).Error

	return options, err
}

func (vr *voteRepository) InsertAnswers(answers []models.UserAnswer) error {
	err := vr.db.Create(&answers).Error
	return err
}

func (vr *voteRepository) GetUserAnswers(voteId, userId string) ([]response.QuestionAnswer, error) {
	var res []response.QuestionAnswer
	err := vr.db.Raw("SELECT a.question_id, b.Question, a.option_id, c.Option FROM user_answers a LEFT JOIN vote_questions b ON a.question_id = b.id LEFT JOIN vote_options c ON a.option_id = c.id WHERE a.vote_id = ? AND a.user_id = ?", voteId, userId).Scan(&res).Error

	return res, err
}
