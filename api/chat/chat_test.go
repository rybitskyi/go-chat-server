package chat

import (
	"testing"
	"github.com/rybitskyi/go-chat-server/api/model"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestChat(t *testing.T) {
	c := New(10)

	msg, err := c.AddMessage(model.Message{User: "superman", Text: "hello1"})
	assert.Nil(t, err)
	assert.NotEmpty(t, msg.Timestamp)

	c.AddMessage(model.Message{User: "batman", Text: "hello2"})

	//Messages
	messages, err := c.GetMessages()
	assert.Nil(t, err)

	assert.Len(t, messages, 2)

	assert.Equal(t, "superman", messages[0].User)
	assert.Equal(t, "hello1", messages[0].Text)
	assert.NotEmpty(t, messages[0].Timestamp)

	assert.Equal(t, "batman", messages[1].User)
	assert.Equal(t, "hello2", messages[1].Text)
	assert.NotEmpty(t, messages[1].Timestamp)

	//Users
	users, err := c.GetUsers()
	assert.Nil(t, err)

	assert.Len(t, users, 2)
	assert.Contains(t, users, "superman")
	assert.Contains(t, users, "batman")
}

//list 100 most recent messages, sorted by 'timestamp' posted to the chat server.
func TestRecentMessages(t *testing.T) {
	c := New(5)

	for i := 0; i < 10; i++ {
		_, err := c.AddMessage(model.Message{User: fmt.Sprintf("user%d", i), Text: fmt.Sprintf("hello%d", i)})
		assert.Nil(t, err)
	}

	//Users
	users, err := c.GetUsers()
	assert.Nil(t, err)
	assert.Len(t, users, 10)

	//Messages
	messages, err := c.GetMessages()
	assert.Nil(t, err)

	assert.Len(t, messages, 5)

	assert.Equal(t, "user5", messages[0].User)
	assert.Equal(t, "hello5", messages[0].Text)

	assert.Equal(t, "user9", messages[4].User)
	assert.Equal(t, "hello9", messages[4].Text)
}

func TestUsers(t *testing.T) {
	c := New(10)

	msg, err := c.AddMessage(model.Message{User: "superman", Text: "hello1"})
	assert.Nil(t, err)
	assert.NotEmpty(t, msg.Timestamp)

	c.AddMessage(model.Message{User: "superman", Text: "hello2"})

	//Users
	users, err := c.GetUsers()
	assert.Nil(t, err)

	assert.Len(t, users, 1)
	assert.Contains(t, users, "superman")
}

/*
TODO: Fix test case
func TestParallelThreads(t *testing.T) {
	c := New(5)
	count := 100000

	for i := 0; i < count; i++ {
		go func(v int) {
			_, err := c.AddMessage(model.Message{User: fmt.Sprintf("user%d", v), Text: fmt.Sprintf("hello%d", v)})
			assert.Nil(t, err)
		}(i)
	}
	time.Sleep(4 * time.Second)

	//Users
	users, err := c.GetUsers()
	assert.Nil(t, err)
	assert.Len(t, users, count)

	//Messages
	messages, err := c.GetMessages()
	assert.Nil(t, err)

	assert.Len(t, messages, 5)
}
*/
