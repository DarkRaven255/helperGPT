package gpt

type Conversation struct {
	Model    string    `json:"model"`
	Scenario string    `json:"-"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func NewMessage(role string, content string) Message {
	return Message{
		Role:    role,
		Content: content,
	}
}
