package main

import (
	"flag"
	"os"
)

func main() {

	prompt := flag.String("prompt", "tell me a joke", "prompt to send to the LLM")
	flag.Parse()

	cb := MakeChatbot(os.Getenv("OPENAI_PROJECT_KEY"))
	cb.Chat(*prompt)
}
