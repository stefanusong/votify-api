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
	GetVotesByUserID(UserID string) (any, error)
	GetVoteQuestions(voteId string) (any, error)
	AnswerVote(req request.UserAnswer) error
	GetUserAnswers(VoteID, UserID string) (any, error)
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
	// TODO: Apply Pagination
	votes, err := svc.voteRepo.GetPublicVotes()
	if err != nil {
		return nil, err
	}

	return map[string][]response.BaseVote{"votes": votes}, nil
}

func (svc *voteService) GetVoteByID(ID string) (any, error) {
	vote, err := svc.voteRepo.GetVoteByID(ID)
	if err != nil {
		return nil, err
	}

	if vote == nil {
		return nil, nil
	}

	return map[string]response.BaseVote{"vote": *vote}, nil
}

func (svc *voteService) GetVotesByUserID(UserID string) (any, error) {
	votes, err := svc.voteRepo.GetVotesByUserID(UserID)
	if err != nil {
		return nil, err
	}

	return map[string][]response.BaseVote{"votes": votes}, nil
}

func (svc *voteService) GetVoteQuestions(voteId string) (any, error) {
	var baseQuestions []response.BaseQuestion

	questions, err := svc.voteRepo.GetVoteQuestions(voteId)
	if err != nil {
		return nil, err
	}

	for _, question := range questions {
		// Get options
		var baseOptions []response.BaseOption
		options, err := svc.voteRepo.GetQuestionOptions(question.ID.String())
		if err != nil {
			return nil, err
		}
		for _, option := range options {
			baseOptions = append(baseOptions, response.BaseOption{ID: option.ID, Option: option.Option, Desc: option.Desc})
		}

		// append questions
		baseQuestions = append(baseQuestions, response.BaseQuestion{ID: question.ID, Question: question.Question, Options: baseOptions})
	}

	return map[string][]response.BaseQuestion{"questions": baseQuestions}, nil
}

func (svc *voteService) AnswerVote(req request.UserAnswer) error {
	var UserAnswers []models.UserAnswer

	// Get vote
	vote, err := svc.voteRepo.GetVoteByID(req.VoteID.String())
	if err != nil {
		return err
	}
	if !vote.IsOpen {
		return errors.New("the vote is already closed")
	}

	// Get questions
	questions, err := svc.voteRepo.GetVoteQuestions(req.VoteID.String())
	if err != nil {
		return err
	}
	for i, question := range questions {
		options, err := svc.voteRepo.GetQuestionOptions(question.ID.String())
		if err != nil {
			return err
		}
		questions[i].VoteOptions = append(question.VoteOptions, options...)
	}

	// Validate and append answers
	for _, answer := range req.Answers {
		validQuestion, validOption := validateQuestionAndOption(questions, answer)

		if !validQuestion {
			return errors.New("invalid question")
		} else if !validOption {
			return errors.New("invalid option")
		}

		UserAnswers = append(UserAnswers, models.NewUserAnswer(req.VoteID, req.UserID, answer.QuestionID, answer.OptionID))
	}

	// Bulk insert answers
	err = svc.voteRepo.InsertAnswers(UserAnswers)
	return err
}

func validateQuestionAndOption(questions []models.VoteQuestion, answer request.Answer) (bool, bool) {
	validQuestion := false
	validOption := false
	for _, q := range questions {
		if q.ID == answer.QuestionID {
			validQuestion = true
			for _, o := range q.VoteOptions {
				if o.ID == answer.OptionID {
					validOption = true
					break
				}
			}
			break
		}
	}
	return validQuestion, validOption
}

func (svc *voteService) GetUserAnswers(VoteID, UserID string) (any, error) {
	res, err := svc.voteRepo.GetUserAnswers(VoteID, UserID)
	if err != nil {
		return nil, err
	}

	userAns := response.UserAnswer{VoteID: VoteID, Questions: res}

	return userAns, nil
}
