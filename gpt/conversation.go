package gpt

func NewConversation(model string, scenario string) *Conversation {
	return &Conversation{
		Model:    model,
		Messages: []Message{},
		Scenario: scenario,
	}
}

func (c *Conversation) AddMessage(message Message) {
	c.Messages = append(c.Messages, message)
}

func (c *Conversation) PrintConversation() string {
	var messages string
	for _, message := range c.Messages[1:] {
		if message.Role != User {
			messages += message.Role + ": " + message.Content + "\n\n"
		} else {
			messages += message.Role + ": " + message.Content + "\n"
		}
	}
	return messages
}
