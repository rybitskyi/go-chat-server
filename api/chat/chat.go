package chat

import (
	"github.com/rybitskyi/go-chat-server/api/model"
	"time"
	"github.com/hashicorp/golang-lru"
	log "github.com/sirupsen/logrus"
	"sync"
)

type ChatService struct {
	messages         *lru.Cache
	messagesCapacity int
	users            sync.Map
}

var exists = struct{}{}

var _ Service = (*ChatService)(nil)

func New(capacity int) *ChatService {
	l, err := lru.New(capacity)
	if err != nil {
		log.Fatal(err)
	}
	c := ChatService{
		messages:         l,
		messagesCapacity: capacity,
	}
	return &c
}

func (c *ChatService) GetMessages() ([]model.Message, error) {
	keys := c.messages.Keys()
	result := make([]model.Message, len(keys))
	for i, k := range keys {
		msg := k.(*model.Message)
		result[i] = *msg
	}
	return result, nil
}

func (c *ChatService) AddMessage(message model.Message) (model.Message, error) {
	//log.Infof("AddMessage %#v", message)
	message.Timestamp = time.Now().Unix()
	c.messages.Add(&message, nil)

	c.users.Store(message.User, exists)

	return message, nil
}

func (c *ChatService) GetUsers() ([]string, error) {
	result := make([]string, 0)
	c.users.Range(func(k, v interface{}) bool {
		result = append(result, k.(string))
		return true
	})
	return result, nil;
}
