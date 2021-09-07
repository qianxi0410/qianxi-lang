package main

import (
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/monkey-lang/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", user.Username)

	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)

	// content, err := ioutil.ReadFile("./examples/arr.ml")
	// if err != nil {
	// 	log.Fatal("err when reading file ")
	// }

	// env := object.NewEnvironment()
	// l := lexer.New(string(content))
	// p := parser.New(l)
	// program := p.ParseProgram()

	// if len(p.Errors()) != 0 {
	// 	printParserErrors(os.Stdout, p.Errors())
	// }

	// evaluated := evaluator.Eval(program, env)

	// if evaluated != nil {
	// 	io.WriteString(os.Stdout, evaluated.Inspect())
	// 	io.WriteString(os.Stdout, "\n")
	// }
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
