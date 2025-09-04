package parser

import (
	"fmt"

	"github.com/alecthomas/participle/v2/lexer"
)

type Option struct {
	Pos lexer.Position

	Name  string  `( "(" @Ident @( "." Ident )* ")" | @Ident @( "." @Ident )* )`
	Attr  *string `( "." @Ident ( "." @Ident )* )?`
	Value *Value  `"=" @@`
}

func (o *Option) Output() string {
	return fmt.Sprintf("option %s = %v", o.Name, o.Value)
}
