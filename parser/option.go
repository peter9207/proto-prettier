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
	if o.Attr != nil {
		fmt.Println("[UNKNOWN]option attr:", *o.Attr)
		return fmt.Sprintf("option %s.%s = [%v]", o.Name, *o.Attr, *o.Value)
	}
	// fmt.Println("option name:", o.Name, *o.Value, o.Attr)
	return fmt.Sprintf("option %s = [%v]", o.Name, o.Value.Output())
}
