package model

type Message struct {
	User      string `json:"user" binding:"required"` // User Name
	Timestamp int64  `json:"timestamp"`               // Timestamp when this message was created
	Text      string `json:"text" binding:"required"` // Chat Message Text
}

type Messages struct {
	Messages []Message `json:"messages"`
}

type Users struct {
	Users []string `json:"users"`
}

type Ok struct {
	OK bool `json:"ok"`
}
