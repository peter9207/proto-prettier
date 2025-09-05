package parser

import (
	"fmt"

	"github.com/alecthomas/participle/v2/lexer"
)

type EnumEntry struct {
	Pos lexer.Position

	Value  *EnumValue  `  @@`
	Option *EnumOption `| "option" @@`
}

type EnumValue struct {
	Pos lexer.Position

	Key   string `@Ident`
	Value int    `"=" @( [ "-" ] Int )`

	Options []*Option `( "[" @@ ( "," @@ )* "]" )?`
}

type EnumOption struct {
	Pos lexer.Position

	Name  string           `( "(" @Ident @( "." Ident )* ")" | @Ident @( "." @Ident )* )`
	Attr  *string          `( "." @Ident ( "." @Ident )* )?`
	Value *EnumOptionValue `"=" @@`
}

type EnumOptionValue struct {
	Pos lexer.Position

	String    *string  `  @String`
	Number    *float64 `| @Float`
	Int       *int64   `| @Int`
	Bool      *bool    `| (@"true" | "false")`
	Reference *string  `| @Ident @( "." Ident )*`
	Map       *Map     `| @@`
	Array     *Array   `| @@`
}

func (ev *EnumOptionValue) Output() string {

	fmt.Println("========enum option value:", ev)

	if ev.String != nil {
		return *ev.String
	}
	if ev.Number != nil {
		return fmt.Sprintf("%f", *ev.Number)
	}
	if ev.Int != nil {
		return fmt.Sprintf("%d", *ev.Int)
	}
	if ev.Bool != nil {
		return fmt.Sprintf("%t", *ev.Bool)
	}

	return "undefined enum option value"
}

func (o *EnumOption) Output() string {
	fmt.Println("========enum option:", o)
	if o.Attr != nil {
		fmt.Println("[UNKNOWN]option attr:", *o.Attr)
		return fmt.Sprintf("option %s.%s = %v", o.Name, *o.Attr, *o.Value)
	}
	// fmt.Println("option name:", o.Name, *o.Value, o.Attr)
	return fmt.Sprintf("option %s = %v", o.Name, o.Value.Output())
}

func (ev *EnumValue) Output() string {

	fmt.Println("========enum value:", ev)

	return fmt.Sprintf("%s = %d;", ev.Key, ev.Value)
}

func (ee *EnumEntry) Output() string {
	fmt.Println("========enum entry:", ee)

	if ee.Value != nil {

		return ee.Value.Output()

	}

	return ee.Option.Output()

}
