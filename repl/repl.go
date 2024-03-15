package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/Linkinlog/MagLang/evaluator"
	"github.com/Linkinlog/MagLang/lexer"
	"github.com/Linkinlog/MagLang/object"
	"github.com/Linkinlog/MagLang/parser"
)

const name = "MagLang"

const PROMPT = "(mag) "

const PUPPEROON = `
     |\_/|                  
     | @ @   Woof? 
     |   <>              _  
     |  _/\------____ ((| |))
     |               ` + "`" + `--' |   
 ____|_       ___|   |___.' 
/_/_____/____/_______|
`

func RunFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("couldnt open file %s\n", filename)
		return
	}

	scanner := bufio.NewScanner(file)
	env := object.NewEnvironment()

	for {
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(os.Stdout, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			fmt.Fprint(os.Stdout, evaluated.Inspect())
			fmt.Fprint(os.Stdout, "\n")
		}
	}
}

func Start(in io.Reader, out io.Writer) {
	greet()

	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			fmt.Fprint(out, "Goodbye! :(\n")
			return
		}
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			fmt.Fprint(out, evaluated.Inspect())
			fmt.Fprint(out, "\n")
		}
	}
}

func greet() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the %s REPL!\n",
		user.Username, name)
	fmt.Printf("Please enter some commands!\n")
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		fmt.Fprint(out, PUPPEROON)
		fmt.Fprint(out, "\tWoof! We ran into a problem here!\n")
		fmt.Fprint(out, "\t parser errors:\n")
		fmt.Fprint(out, "\t"+msg+"\n")
	}
}
