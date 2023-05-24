package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Linkinlog/MagLang/lexer"
	"github.com/Linkinlog/MagLang/token"
)

const PROMPT = "(mag-repl) "

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
		for toke := l.NextToken(); toke.Type != token.EOF; toke = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", toke)
		}
	}
}
