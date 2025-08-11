package repl

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/batt0s/rizzy/evaluator"
	"github.com/batt0s/rizzy/lexer"
	"github.com/batt0s/rizzy/object"
	"github.com/batt0s/rizzy/parser"
	"github.com/chzyer/readline"
)

const PROMPT = ">>> "
const HistoryFile = ".rizzy_history"

func Start(in io.Reader, out io.Writer) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          PROMPT,
		HistoryFile:     HistoryFile,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
		AutoComplete:    completer{},
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	env := object.NewEnvironment()

	var lines []string
	var openBrackets int

	for {
		var line string
		if openBrackets > 0 {
			rl.SetPrompt("...> ")
		} else {
			rl.SetPrompt(PROMPT)
		}

		input, err := rl.Readline()
		if err == readline.ErrInterrupt {
			if len(lines) == 0 {
				break
			} else {
				lines = nil
				openBrackets = 0
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(input)
		if line == "" && openBrackets == 0 {
			continue
		}

		lines = append(lines, line)
		openBrackets += strings.Count(line, "{")
		openBrackets -= strings.Count(line, "}")

		if openBrackets > 0 {
			continue
		}

		fullInput := strings.Join(lines, " ")

		lines = nil
		openBrackets = 0

		l := lexer.New(fullInput)
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

type completer struct{}

var keywords = []string{
	"func",
	"def",
	"true",
	"false",
	"if",
	"else",
	"return",
	// Basics
	"type",
	"puts",
	"rizz",
	"fmt",
	"exit",
	// Array Operations
	"len",
	"first",
	"last",
	"head",
	"tail",
	"push",
	"pop",
	"range",
	// Math
	"pow",
	"sqrt",
	// Types
	"int",
	"float",
}

func (c completer) Do(line []rune, pos int) ([][]rune, int) {
	input := string(line[:pos])
	fields := strings.Fields(input)

	var prefix string
	if len(fields) > 0 {
		prefix = fields[len(fields)-1]
	} else {
		prefix = ""
	}

	var suggestions [][]rune
	for _, kw := range keywords {
		if strings.HasPrefix(kw, prefix) {
			suggestions = append(suggestions, []rune(kw[len(prefix):]))
		}
	}
	return suggestions, len(prefix)
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
