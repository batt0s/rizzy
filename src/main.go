package main

import (
	"errors"
	"fmt"
	"os"
	"os/user"

	"github.com/batt0s/rizzy/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	cliArgs := os.Args
	if len(cliArgs) > 1 {
		filePath := cliArgs[1]
		if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Couldn't find file: %s\n", filePath)
			os.Exit(1)
		}
		repl.RunFile(filePath, os.Stdout)
	} else {
		fmt.Printf("Hello %s! This is the Rizzler!\n", user.Username)
		repl.Start(os.Stdin, os.Stdout)
	}
}
