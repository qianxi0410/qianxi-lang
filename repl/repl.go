package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/monkey-lang/evaluator"
	"github.com/monkey-lang/lexer"
	"github.com/monkey-lang/object"
	"github.com/monkey-lang/parser"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf("%s ", PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
