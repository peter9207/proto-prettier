package parser

import (
	"strings"

	"github.com/alecthomas/participle/v2/lexer"
)

type MessageEntry struct {
	Pos lexer.Position

	Enum       *Enum       `( @@`
	Option     *Option     ` | "option" @@`
	Message    *Message    ` | @@`
	Oneof      *Oneof      ` | @@`
	Extend     *Extend     ` | @@`
	Reserved   *Reserved   ` | @@`
	Extensions *Extensions ` | @@`
	Field      *Field      ` | @@ ) ";"*`
}

func (me *MessageEntry) Output() string {

	entries := []string{}

	if me.Enum != nil {
		entries = append(entries, me.Enum.Output())
	}

	if me.Option != nil {
		entries = append(entries, "undefined option")
	}

	if me.Field != nil {
		entries = append(entries, me.Field.Output())
	}
	return strings.Join(entries, " ")
}
