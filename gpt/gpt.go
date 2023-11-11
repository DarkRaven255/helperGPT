package gpt

import (
	"helperGPT/gpt/scenario"
)

type GptScenario struct {
	Temperature  float32       `json:"temperature,omitempty"`
	Role         string        `json:"role,omitempty"`
	Model        string        `json:"model,omitempty"`
	Content      string        `json:"content,omitempty"`
	Conversation *Conversation `json:"-"`
}

const (
	System    = "system"
	Assistant = "assistant"
	User      = "user"
	Function  = "function"
)

const (
	GPT3_5Turbo = "gpt-3.5-turbo"
	GPT4        = "gpt-4"
)

func NewGpt(model string, scenario string) *GptScenario {
	gptScenario := GptScenario{}

	gptScenario.Role = System
	gptScenario.Model = model

	conversation := NewConversation(gptScenario.Model, scenario)
	conversation.AddMessage(NewMessage(gptScenario.Role, scenarioDecoder(scenario)))

	gptScenario.Conversation = conversation

	return &gptScenario

}

func scenarioDecoder(scenarioName string) string {
	switch scenarioName {
	case scenario.Chatbot:
		return Chatbot
	case scenario.Programmer:
		return Programmer
	case scenario.Teacher:
		return Teacher
	case scenario.PolishEnglishTranslator:
		return PolishEnglishTranslator
	default:
		return PersonalAssistant
	}
}

const (
	PersonalAssistant       = "You are a helpful personal assistant."
	Chatbot                 = "You are a helpful chatbot."
	Programmer              = "You are a helpful programmer."
	Teacher                 = "You are a helpful teacher."
	PolishEnglishTranslator = "You are a helpful Polish/English translator. Your role is to translate Polish to English and English to Polish. You are not responding to the questions, you are just translating."
)
