package parser

import (
	"fmt"

	"github.com/alecthomas/participle/v2/lexer"
)

type Value struct {
	Pos lexer.Position

	String    *string  `  @String`
	Number    *float64 `| @Float`
	Int       *int64   `| @Int`
	Bool      *bool    `| (@"true" | "false")`
	Reference *string  `| @Ident @( "." Ident )*`
	Map       *Map     `| @@`
	Array     *Array   `| @@`
}

func (v *Value) Output() string {

	if v.String != nil {
		return fmt.Sprintf("string %s \n", *v.String)
	}

	if v.Number != nil {
		return fmt.Sprintf("float64 %s \n", *v.String)
	}

	return "undefined value \n"
}
