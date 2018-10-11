package chat

import "github.com/rybitskyi/go-chat-server/api/model"

type Service interface {
	GetMessages() ([]model.Message, error)

	AddMessage(message model.Message) (model.Message, error)

	GetUsers() ([]string, error)
}
