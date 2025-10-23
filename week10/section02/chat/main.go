package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	apiKey := os.Getenv("OPENAI_PROJECT_KEY")
	client := openai.NewClient(apiKey)

	prompt := "tell me a joke"

	req := openai.ChatCompletionRequest{
		Model: openai.GPT5,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}

	resp, err := client.CreateChatCompletion(context.TODO(), req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
