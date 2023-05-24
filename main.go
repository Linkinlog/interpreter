package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Linkinlog/MagLang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the Maglang REPL!\n",
		user.Username)
	fmt.Printf("Please enter some commands!\n")
	repl.Start(os.Stdin, os.Stdout)
}
