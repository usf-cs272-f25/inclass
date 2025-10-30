package main

import (
	"bufio"
	"fmt"
	"os"
)

func Chat(t string) {
	// chatbot works goes here
	fmt.Println(t)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What? ")

	// Scan() runs until CTRL-D
	for scanner.Scan() {
		t := scanner.Text()
		Chat(t)
		fmt.Print("What? ")
	}
}
