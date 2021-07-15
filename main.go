package main

import (
	"log"
	"os"

	"gt/commit"
	gtInit "gt/init"
)

func main() {

	command := os.Args[1]

	switch command {
	case "init":
		gtInit.Init()
	case "commit":
		commit.Commit()
	default:
		log.Printf("gt : '%s' is not a gt command ", command)
	}
}
