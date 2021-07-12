package main

import (
	"log"
	"os"

	"gt/commit"
	"gt/init"
)

func main() {

	command := os.Args[1]

	switch command {
	case "init":
		init.Init()
	case "commit":
		commit.Commit()
	default:
		log.Printf("gt : '%s' is not a gt command ", command)
	}
}
