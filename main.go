package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to Monkeylang!\n", user.Username)
	fmt.Printf("Feel free to start typing in commands :)\n")
	repl.Start(os.Stdin, os.Stdout)
}
