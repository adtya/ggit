package main

import (
	"fmt"
	"log"
	"os"

	"adtya.xyz/ggit/commands"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: ggit <command>")
		os.Exit(1)
	}

	var str string
	var err error

	switch args[0] {
	case "init":
		str = commands.Init(args[1:])
	case "hash-object":
		str, err = commands.HashObject(args[1:])
	case "cat-file":
		str, err = commands.CatFile(args[1:])
	default:
		fmt.Printf("%s is not a valid command\n", args[0])
		os.Exit(1)
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(str)
}
