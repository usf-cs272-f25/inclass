package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	// Create a file
	fileName := "example.txt"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	// Write to the file
	_, err = file.WriteString("Hello, World!\n")
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	// Read the file
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Println(string(data))

	// Delete the file
	err = os.Remove(fileName)
	if err != nil {
		log.Fatalf("Error deleting file: %v", err)
	}
}