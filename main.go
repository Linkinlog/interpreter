package main

import (
	"os"

	"github.com/Linkinlog/MagLang/repl"
)

func main() {
	if len(os.Args) > 1 {
		repl.RunFile(os.Args[1])
		return
	}
	repl.Start(os.Stdin, os.Stdout)
}
