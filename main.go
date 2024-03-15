package main

import (
	"os"

	"github.com/Linkinlog/MagLang/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
