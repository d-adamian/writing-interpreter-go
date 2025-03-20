package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/d-adamian/writing-interpreter-go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Print("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

// Continue at page 41, after "we can parse let statement!"
