package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/tsingbx/compiler/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the %s programming language!\n", user.Username, repl.LANGUAGE)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
