package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

func GetCurrentWeather(tc openai.ToolCall) string {
	args := map[string]interface{}{}
	err := json.Unmarshal([]byte(tc.Function.Arguments), &args)
	if err != nil {
		log.Fatalln("json.Unmarshal: ", err)
	}

	// Use a Go type assertion to cast the interface{} to string
	// or "instructor": "Phil Peterson"
	location := args["location"].(string)
	switch tc.Function.Name {
	// get_courses_for_instructor
	case "get_current_weather":
		if location == "San Francisco, CA" {
			return "55 and foggy"
		}
	default:
		log.Fatal("unknown function name")
	}
	return ""
}

func main() {
	// Use JSONSchema to describe the parameters and their types
	params := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"location": {
				Type:        jsonschema.String,
				Description: "The location where we want the weather forcast",
			},
		},
		Required: []string{"location"},
	}

	// Make a function using those parameters
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

	// Use a slice of ChatCompletionMessage to maintain context with the LLM
	dialogue := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "What is the weather in San Francisco?",
		},
	}

	req := openai.ChatCompletionRequest{
		Model:    openai.GPT5,
		Messages: dialogue,
		Tools:    []openai.Tool{t},
	}

	client := openai.NewClient(os.Getenv("OPENAI_PROJECT_KEY"))
	resp, err := client.CreateChatCompletion(context.TODO(), req)
	if err != nil {
		log.Fatal("CreateChatCompletion: ", err)
	}
	msg := resp.Choices[0].Message

	if len(msg.ToolCalls) != 1 {
		log.Fatal("expected one tool call")
	}

	dialogue = append(dialogue, msg)

	newmsg := openai.ChatCompletionMessage{
		Role:       openai.ChatMessageRoleTool,
		Content:    GetCurrentWeather(msg.ToolCalls[0]),
		Name:       msg.ToolCalls[0].Function.Name,
		ToolCallID: msg.ToolCalls[0].ID,
	}

	dialogue = append(dialogue, newmsg)

	newreq := openai.ChatCompletionRequest{
		Model:    openai.GPT5,
		Messages: dialogue,
		Tools:    []openai.Tool{t},
	}
	resp, err = client.CreateChatCompletion(context.TODO(), newreq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Choices[0].Message.Content)
}
