package main

import (
	"log"
	"os"
	"path"

	"gt/commit"
)

// this function for init gt
// if success it's will be create folder .git
// commit for run init
// gt init <filename>
func Init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}

	foldernames := os.Args
	var foldername string
	if len(foldernames) > 2 {
		foldername = foldernames[2]
	}

	if foldername != "" && foldername != "." && foldername != ".." {
		pwd = path.Join(pwd, foldername)
	}

	git_path := path.Join(pwd, ".git")
	if err := os.Mkdir(git_path, 0755); err != nil {
		log.Println(err)
		return
	}

	var folders = []string{"objects", "refs"}

	for _, folder := range folders {
		if err := os.Mkdir(path.Join(git_path, folder), 0755); err != nil {
			log.Printf("%s: %v\n", git_path, err)
			return
		}
	}

	log.Printf("Initialize empty gt repository in %s", git_path)
}

func main() {

	command := os.Args[1]

	switch command {
	case "init":
		Init()
	case "commit":
		commit.Commit()
	default:
		log.Printf("gt : '%s' is not a gt command ", command)
	}
}
