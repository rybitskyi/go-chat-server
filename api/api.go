package api

import (
	"github.com/rybitskyi/go-chat-server/api/chat"
	"github.com/gin-gonic/gin"
)

const (
	ChatMessagesCapacity = 100
)

type APIService struct {
	route       *gin.Engine
	ChatService *chat.ChatService
}

func New(r *gin.Engine) *APIService {
	api := &APIService{
		route:       r,
		ChatService: chat.New(ChatMessagesCapacity),
	}
	api.initRoutesAndEndpoints()
	return api
}

func (s *APIService) initRoutesAndEndpoints() {
	s.route.GET("/status", getStatus)

	chat.AddRoutes(s.route, s.ChatService)
}

func getStatus(c *gin.Context) {
	c.JSON(200, "alive")
}
