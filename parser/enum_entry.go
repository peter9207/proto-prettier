package parser

import (
	"fmt"

	"github.com/alecthomas/participle/v2/lexer"
)

type EnumEntry struct {
	Pos lexer.Position

	Value  *EnumValue `  @@`
	Option *Option    `| "option" @@`
}

type EnumValue struct {
	Pos lexer.Position

	Key   string `@Ident`
	Value int    `"=" @( [ "-" ] Int )`

	Options []*Option `( "[" @@ ( "," @@ )* "]" )?`
}

func (ev *EnumValue) Output() string {

	return fmt.Sprintf("%s = %d;", ev.Key, ev.Value)
}

func (ee *EnumEntry) Output() string {

	if ee.Value != nil {

		return ee.Value.Output()

	}

	return ee.Option.Output()

}
