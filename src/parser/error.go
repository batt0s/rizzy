package parser

import "fmt"

type Error struct {
	msg  string
	line int
	col  int
}

func (e *Error) String() string {
	return fmt.Sprintf("Parser Error on line %d, col %d: %s", e.line, e.col, e.msg)
}
