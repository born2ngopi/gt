package main

import (
	"log"
	"os"

	"github.com/needkopi/gt/commit"
	"github.com/needkopi/gt/init"
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
