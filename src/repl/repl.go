package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/batt0s/rizzy/evaluator"
	"github.com/batt0s/rizzy/lexer"
	"github.com/batt0s/rizzy/object"
	"github.com/batt0s/rizzy/parser"
)

const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)

		var lines []string
		brackets := 0
		for {
			scanned := scanner.Scan()
			if !scanned {
				return
			}
			line := scanner.Text()
			lines = append(lines, line)
			brackets += strings.Count(line, "{")
			brackets -= strings.Count(line, "}")
			if brackets == 0 {
				break
			}

			fmt.Print("...>")
		}

		input := strings.Join(lines, " ")

		l := lexer.New(input)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, "Rizzler: ")
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func RunFile(filepath string, out io.Writer) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	var builder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		builder.WriteString(scanner.Text() + " ")
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	input := builder.String()

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return nil
	}

	evaluated := evaluator.Eval(program, object.NewEnvironment())
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect()+"\n")
	}

	return nil
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, msg+"\n")
	}
}
