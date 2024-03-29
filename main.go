package main

import (
	"fmt"
	"github.com/snoozer05/monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel freeto type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
