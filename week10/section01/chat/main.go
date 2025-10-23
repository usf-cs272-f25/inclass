package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

/* Example of using Go struct tags to translate attrs to JSON
type Foo struct {
	Bar int `json:"bar"`
	SomeThing int `json:"some_thing`
}
*/

func main() {
	// OpenAI's APIs are simply an HTTP endpoint
	// with a JSON payload
	// The go-openai package wraps those JSON payloads using
	// Go structs with `json:` struct tags

	// use an env var to store the API key so you don't
	// check it in, which is a security issue
	apiKey := os.Getenv("OPENAI_PROJECT_KEY")
	client := openai.NewClient(apiKey)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT5,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "tell me a joke",
			},
		},
	}

	resp, err := client.CreateChatCompletion(context.TODO(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Choices[0].Message.Content)
}
