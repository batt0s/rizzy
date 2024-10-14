package evaluator

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/batt0s/rizzy/object"
)

var builtins = map[string]*object.Builtin{
	// Basics
	"type": &object.Builtin{Fn: builtin_type},
	"puts": &object.Builtin{Fn: builtin_puts},
	"rizz": &object.Builtin{Fn: builtin_puts},
	"fmt":  &object.Builtin{Fn: builtin_fmt},
	"exit": &object.Builtin{Fn: builtin_exit},
	// Array Operations
	"len":   &object.Builtin{Fn: builtin_len},
	"first": &object.Builtin{Fn: builtin_first},
	"last":  &object.Builtin{Fn: builtin_last},
	"head":  &object.Builtin{Fn: builtin_head},
	"tail":  &object.Builtin{Fn: builtin_tail},
	"push":  &object.Builtin{Fn: builtin_push},
	"pop":   &object.Builtin{Fn: builtin_pop},
	// Math
	"pow":  &object.Builtin{Fn: builtin_pow},
	"sqrt": &object.Builtin{Fn: builtin_sqrt},
}

func builtin_type(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	return &object.String{Value: string(args[0].Type())}
}

func builtin_puts(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}

	return NULL
}

func builtin_fmt(args ...object.Object) object.Object {
	if len(args) < 2 {
		return newError("wrong number of arguments. got=%d, want=2", len(args))
	}
	if args[0].Type() != object.STRING_OBJ {
		return newError("argument to `fmt` must be STRING, got %s",
			args[0].Type())
	}
	s := args[0].(*object.String).Value
	format := strings.Count(s, "%%")
	if format != len(args)-1 {
		return newError("wrong number of arguments. got=%d, want=%d",
			len(args), format+1)
	}
	for _, arg := range args[1:] {
		s = strings.Replace(s, "%%", arg.Inspect(), 1)
	}
	return &object.String{Value: s}
}

func builtin_exit(args ...object.Object) object.Object {
	code := 0
	if len(args) == 1 && args[0].Type() == object.INTEGER_OBJ {
		code = int(args[0].(*object.Integer).Value)
	}
	os.Exit(code)

	return NULL
}

func builtin_len(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	default:
		return newError("argument to `len` not supported, got %s", args[0].Type())
	}
}

func builtin_first(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `first` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return NULL
}

func builtin_last(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `last` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[len(arr.Elements)-1]
	}

	return NULL
}

func builtin_head(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `head` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if len(arr.Elements) > 0 {
		newElements := make([]object.Object, length-1, length-1)
		copy(newElements, arr.Elements[:length-1])
		return &object.Array{Elements: newElements}
	}

	return NULL
}

func builtin_tail(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `tail` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if len(arr.Elements) > 0 {
		newElements := make([]object.Object, length-1, length-1)
		copy(newElements, arr.Elements[1:length])
		return &object.Array{Elements: newElements}
	}

	return NULL
}

func builtin_push(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=2",
			len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `push` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)

	newElements := make([]object.Object, length+1, length+1)
	copy(newElements, arr.Elements)
	newElements[length] = args[1]

	return &object.Array{Elements: newElements}
}

func builtin_pop(args ...object.Object) object.Object {
	if len(args) < 1 || len(args) > 2 {
		return newError("wrong number of arguments. got=%d, want=1 or 2",
			len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `pop` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length <= 0 {
		return arr
	}

	idx := length - 1
	if len(args) == 2 {
		if args[1].Type() != object.INTEGER_OBJ {
			return newError("argument to `pop` must be INTEGER, got %s",
				args[1].Type())
		}
		idx = int(args[1].(*object.Integer).Value)
		if idx > length-1 {
			return arr
		}
	}

	newElements := make([]object.Object, length-1, length-1)
	copy(newElements, append(arr.Elements[:idx], arr.Elements[idx+1:]...))

	return &object.Array{Elements: newElements}
}

// Math
func builtin_pow(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=2",
			len(args))
	}

	if args[0].Type() != object.INTEGER_OBJ || args[1].Type() != object.INTEGER_OBJ {
		return newError("argument to `pow` must be INTEGER, got %s and %s",
			args[0].Type(), args[1].Type())
	}

	n := float64(args[0].(*object.Integer).Value)
	pow := float64(args[1].(*object.Integer).Value)

	if pow < 0 {
		return newError("argument to `pow` must be positive, got %s",
			args[1].Inspect())
	}

	result := math.Pow(n, pow)
	resultObj := &object.Integer{Value: int64(result)}
	return resultObj
}

func builtin_sqrt(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}

	if args[0].Type() != object.INTEGER_OBJ {
		return newError("argument to `sqrt` must be INTEGER, got %s",
			args[0].Type())
	}

	n := float64(args[0].(*object.Integer).Value)

	if n < 0 {
		return newError("argument to `sqrt` must be positive, got %s",
			args[0].Inspect())
	}

	result := math.Sqrt(n)
	resultObj := &object.Integer{Value: int64(result)}
	return resultObj
}
