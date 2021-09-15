package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/monkey-lang/file"
	"github.com/monkey-lang/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	if len(os.Args) != 1 {
		filePath := os.Args[1]

		file.Start(filePath, os.Stdout)

	} else {
		fmt.Printf("Hello %s!\n", user.Username)

		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}

}
