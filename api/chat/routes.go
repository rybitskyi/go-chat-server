package chat

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/rybitskyi/go-chat-server/api/model"
	log "github.com/sirupsen/logrus"
)

func AddRoutes(r *gin.Engine, chatService *ChatService) {
	r.GET("/messages", getMessages(chatService))
	r.POST("/message", addMessage(chatService))
	r.GET("/users", getUsers(chatService))
}

//list 100 most recent messages, sorted by 'timestamp' posted to the chat server.
func getMessages(chatService *ChatService) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		log.Infof("getMessages handler")
		messages, err := chatService.GetMessages()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		}
		result := model.Messages{Messages: messages}
		c.JSON(http.StatusOK, result)
	}
	return gin.HandlerFunc(fn)
}

//a request to post the given message. when the message is processed by the server a unix timestamp is recorded with each message.
func addMessage(chatService *ChatService) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		log.Infof("addMessage handler")
		var msg model.Message

		if err := c.BindJSON(&msg); err != nil {
			c.JSON(http.StatusPreconditionFailed, err.Error())
			return
		}
		if _, err := chatService.AddMessage(msg); err != nil {
			log.Error("addMessage error", err)
			//c.JSON(http.StatusInternalServerError, err.Error())
			c.JSON(http.StatusOK, model.Ok{false})
		} else {
			c.JSON(http.StatusOK, model.Ok{true})
		}
	}
	return gin.HandlerFunc(fn)
}

//a request to return a set of users that have posted messages so far.
func getUsers(chatService *ChatService) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		log.Infof("getUsers handler")
		users, err := chatService.GetUsers()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		}
		result := model.Users{Users: users}
		c.JSON(http.StatusOK, result)
	}
	return gin.HandlerFunc(fn)
}
