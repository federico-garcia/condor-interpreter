package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/federico-garcia/condor-interpreter/lexer"
	"github.com/federico-garcia/condor-interpreter/token"
)

const PROMPT = ">> "

// Start innitiates the Condor REPL (interactive mode)
func Start(in io.Reader, our io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		tok := l.NextToken()
		for tok.Type != token.EOF {
			fmt.Printf("%+v\n", tok)
			tok = l.NextToken()
		}

	}
}
