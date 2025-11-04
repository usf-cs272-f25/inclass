package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

func get_current_weather(args string) string {
	m := make(map[string]any)
	err := json.Unmarshal([]byte(args), &m)
	if err != nil {
		log.Fatal("unmarshal JSON: ", err)
	}
	for k, v := range m {
		if k == "location" && strings.Contains(v.(string), "Boston") {
			return "Rainy and 55 degrees"
		}

	}
	return ""
}

func main() {
	client := openai.NewClient(os.Getenv("OPENAI_PROJECT_KEY"))

	// Make function params using JSONSchema
	params := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"location": {
				Type:        jsonschema.String,
				Description: "The location where we want the weather forecast",
			},
			"unit": {
				Type: jsonschema.String,
				Enum: []string{"celsius", "fahrenheit"},
			},
		},
		Required: []string{"location"},
	}

	// Make a function using params
	f := openai.FunctionDefinition{
		Name:        "get_current_weather",
		Description: "The city and state, e.g. San Francisco, CA",
		Parameters:  params,
	}

	// Make a tool using the function
	t := openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: &f,
	}

	// Notice this is a slice, so we can have a conversation
	dialogue := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "What is the weather in Boston today?",
		},
	}

	req := openai.ChatCompletionRequest{
		Model:    openai.GPT5,
		Messages: dialogue,
		Tools:    []openai.Tool{t},
	}
	resp, err := client.CreateChatCompletion(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}
	msg := resp.Choices[0].Message
	if len(msg.ToolCalls) != 1 {
		log.Fatal("expected one tool call")
	}

	// Append the LLM's response (and its ID) to the dialogue
	dialogue = append(dialogue, msg)

	newmsg := openai.ChatCompletionMessage{
		Role:       openai.ChatMessageRoleTool,
		Content:    get_current_weather(msg.ToolCalls[0].Function.Arguments),
		Name:       msg.ToolCalls[0].Function.Name,
		ToolCallID: msg.ToolCalls[0].ID,
	}

	// Append our tool call response to the dialogue
	dialogue = append(dialogue, newmsg)
	newreq := openai.ChatCompletionRequest{
		Model:    openai.GPT5,
		Messages: dialogue,
		Tools:    []openai.Tool{t},
	}
	resp, err = client.CreateChatCompletion(context.TODO(), newreq)
	if err != nil {
		log.Fatal("tool call response: ", err)
	}

	log.Println(resp.Choices[0].Message.Content)
}
