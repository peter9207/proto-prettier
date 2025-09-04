package parser

import (
	"fmt"
	// "io"

	// "github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)


type Enum struct {
	Pos lexer.Position

	Name   string       `"enum" @Ident`
	Values []*EnumEntry `"{" ( @@ ( ";" )* )* "}"`
}


func (e *Enum) Output() string {

	start:= fmt.Sprintf(`enum %s`, e.Name)
	return start
}