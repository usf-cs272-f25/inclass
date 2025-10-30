package main

import (
	"bufio"
	"fmt"
	"os"
)

func Chat(t string) {
	// Chatbot work goes here
	fmt.Println(t)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Runs until EOF (CTRL-D)
	fmt.Print("Course catalog> ")
	for scanner.Scan() {
		t := scanner.Text()
		Chat(t)
		fmt.Print("Course catalog> ")
	}
}
