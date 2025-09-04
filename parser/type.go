package parser

import (
	"fmt"

	"github.com/alecthomas/participle/v2/lexer"
)

type Type struct {
	Pos lexer.Position

	Scalar    Scalar   `  @@`
	Map       *MapType `| @@`
	Reference string   `| @(Ident ( "." Ident )*)`
}

func (t *Type) Output() string {
	return fmt.Sprintf("undefined Type \n")
}
