package services

import (
	"errors"
	"regexp"

	"github.com/stefanusong/votify-api/dto/request"
	"github.com/stefanusong/votify-api/dto/response"
	"github.com/stefanusong/votify-api/models"
	"github.com/stefanusong/votify-api/repositories"
)

type VoteService interface {
	CreateVote(voteReq request.CreateVote) (any, error)
	GetPublicVotes() (any, error)
	GetVoteByID(ID string) (any, error)
}

type voteService struct {
	voteRepo repositories.VoteRepository
}

func NewVoteService(voteRepo repositories.VoteRepository) VoteService {
	return &voteService{
		voteRepo: voteRepo,
	}
}

func (svc *voteService) CreateVote(voteReq request.CreateVote) (any, error) {

	isSlugValid := regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`).MatchString(voteReq.Slug)
	if !isSlugValid {
		return nil, errors.New("invalid slug format")
	}

	voteQuestions := bindQuestionsToModel(voteReq.Questions)
	newVote := models.NewVote(voteReq.UserID, voteReq.CategoryID, voteReq.Slug, voteReq.Title,
		voteReq.IsOpen, voteReq.IsPublic, voteQuestions)

	voteId, err := svc.voteRepo.InsertVote(newVote)
	if err != nil {
		return nil, err
	}

	return map[string]string{"vote_id": voteId.String()}, nil
}

func bindQuestionsToModel(questionReq []request.CreateQuestion) []models.VoteQuestion {
	questions := make([]models.VoteQuestion, 0)

	for _, question := range questionReq {
		questionOptions := bindOptionsToModel(question.Options)
		question := models.NewVoteQuestion(question.Question, questionOptions)
		questions = append(questions, question)
	}

	return questions
}

func bindOptionsToModel(optionReq []request.CreateOption) []models.VoteOption {
	options := make([]models.VoteOption, 0)

	for _, option := range optionReq {
		voteOption := models.NewVoteOption(option.Option, option.Desc)
		options = append(options, voteOption)
	}

	return options
}

func (svc *voteService) GetPublicVotes() (any, error) {
	votes, err := svc.voteRepo.GetPublicVotes()

	resp := []response.BaseVote{}
	for _, vote := range votes {
		resp = append(resp, response.NewBaseVote(vote))
	}

	if err != nil {
		return nil, err
	}

	return map[string][]response.BaseVote{"votes": resp}, nil
}

func (svc *voteService) GetVoteByID(ID string) (any, error) {
	vote, err := svc.voteRepo.GetVoteByID(ID)

	if err != nil {
		return nil, err
	}

	if vote == nil {
		return nil, nil
	}

	return map[string]any{"vote": response.NewBaseVote(*vote)}, nil
}
