package main

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

type Chatbot struct {
	client *openai.Client
}

func MakeChatbot(key string) *Chatbot {
	client := openai.NewClient(key)
	return &Chatbot{
		client: client,
	}
}

func (bot *Chatbot) Chat(prompt string) {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	}
	req := openai.ChatCompletionRequest{
		Model:    openai.GPT5,
		Messages: messages,
	}

	resp, err := bot.client.CreateChatCompletion(context.TODO(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Choices[0].Message.Content)
}
