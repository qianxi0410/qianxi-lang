package file

import (
	"io"
	"io/ioutil"
	"log"

	"github.com/monkey-lang/evaluator"
	"github.com/monkey-lang/lexer"
	"github.com/monkey-lang/object"
	"github.com/monkey-lang/parser"
)

func Start(filePath string, out io.Writer) {
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatalf("can not find file: %s", filePath)
	}

	env := object.NewEnvironment()
	l := lexer.New(string(content))
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
	}

	evaluated := evaluator.Eval(program, env)

	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
