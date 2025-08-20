package main

import "fmt"

func printit(s string) (string, error) {
	fmt.Println(s)
	return s, nil
}

func main() {
	_, err := printit("Hello, World")
	if err != nil {
		fmt.Println("how can you fail to print????")
	}
}
