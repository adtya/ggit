package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Usage: %s <command>\n", os.Args[0])
		os.Exit(1)
	}
	switch args[0] {
	case "init":
		Init(args[1:])
	default:
		fmt.Printf("%s is an invalid command\n", args[0])
		os.Exit(1)
	}
}

func Init(args []string) {
	fmt.Println("Hello, World!")
}
