package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	passedInArgs := args[1:]

	if len(passedInArgs) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(passedInArgs) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := passedInArgs[0]
	fmt.Printf("starting crawl %v\n", baseURL)
}
