package handlers

import (
	"github.com/Defenestrationq/forum-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	// Get
	GetAllMessage(c *gin.Context)
	GetMessageByUserID(c *gin.Context)
	GetAllDiscussion(c *gin.Context)
	GetDiscussionByID(c *gin.Context)
	GetDiscussionsByUserID(c *gin.Context)
	// Post
	CreateMessage(c *gin.Context)
	CreateDiscussions(c *gin.Context)
	// Delete
	UpdateMessage(c *gin.Context)
	UpdateDiscussions(c *gin.Context)
	// Update
	DeleteMessage(c *gin.Context)
	DeleteDiscussions(c *gin.Context)
}

type ForumHandler struct {
	dis usecase.DisUseCase
	msg usecase.MsgUseCase
}

func NewHandler(dis usecase.DisUseCase, msg usecase.MsgUseCase) Handler {
	return &ForumHandler{
		dis: dis,
		msg: msg,
	}
}
