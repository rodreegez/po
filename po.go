package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting current working directory")
	}

	user, err := user.Current()
	if err != nil {
		fmt.Println("Error getting user")
	}

	targetSymlink := user.HomeDir + "/.pow/" + filepath.Base(cwd)

	if flag.Arg(0) == "link" {
		err := os.Symlink(cwd, targetSymlink)
		if err != nil {
			fmt.Println("There was an error creating your link")
		} else {
			fmt.Printf("linked %s to %s\n", cwd, targetSymlink)
		}
	} else {
		fmt.Println("Y U NO LINK?")
	}
}
