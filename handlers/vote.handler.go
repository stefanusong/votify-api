package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stefanusong/votify-api/dto/request"
	"github.com/stefanusong/votify-api/dto/response"
	"github.com/stefanusong/votify-api/services"
)

type VoteHandler interface {
	CreateVote(c *gin.Context)
	GetPublicVotes(c *gin.Context)
	GetVoteByID(c *gin.Context)
}

type voteHandler struct {
	voteService services.VoteService
}

func NewVoteHandler(voteService services.VoteService) VoteHandler {
	return &voteHandler{
		voteService: voteService,
	}
}

func (handler *voteHandler) CreateVote(c *gin.Context) {
	// Bind vote
	var newVote request.CreateVote
	if err := c.ShouldBindJSON(&newVote); err != nil {
		resp := response.New(false, "Failed to create new vote", nil, err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	userId := c.GetString("userid")
	newVote.UserID, _ = uuid.FromString(userId)

	// Create vote
	data, err := handler.voteService.CreateVote(newVote)
	if err != nil {
		resp := response.New(false, "Failed to create new vote", nil, err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := response.New(true, "Vote created", data, nil)
	c.JSON(http.StatusCreated, resp)
}

func (handler *voteHandler) GetPublicVotes(c *gin.Context) {
	votes, err := handler.voteService.GetPublicVotes()
	if err != nil {
		resp := response.New(false, "Failed to get public votes", nil, err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := response.New(true, "Success", votes, nil)
	c.JSON(http.StatusOK, resp)
}

func (handler *voteHandler) GetVoteByID(c *gin.Context) {
	voteId := c.Param("id")
	vote, err := handler.voteService.GetVoteByID(voteId)

	if err != nil {
		resp := response.New(false, "Failed to get vote", nil, err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if vote == nil {
		resp := response.New(false, "Failed to get vote", nil, "vote not found")
		c.JSON(http.StatusNotFound, resp)
		return
	}

	resp := response.New(true, "Success", vote, nil)
	c.JSON(http.StatusOK, resp)
}
