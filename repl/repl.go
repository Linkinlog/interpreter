package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Linkinlog/MagLang/evaluator"
	"github.com/Linkinlog/MagLang/lexer"
	"github.com/Linkinlog/MagLang/parser"
)

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

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			fmt.Fprint(out, evaluated.Inspect())
			fmt.Fprint(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		fmt.Fprint(out, PUPPEROON)
		fmt.Fprint(out, "\tWoof! We ran into a problem here!\n")
		fmt.Fprint(out, "\t parser errors:\n")
		fmt.Fprint(out, "\t"+msg+"\n")
	}
}
