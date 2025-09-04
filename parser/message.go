package parser

import (
	"fmt"
	"strings"

	// "io"

	// "github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Message struct {
	Pos lexer.Position

	Name    string          `"message" @Ident`
	Entries []*MessageEntry `"{" @@* "}"`
}

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

	if me.Enum != nil {
		return me.Enum.Output()
	}

	if me.Option != nil {
		return "undefined option"
	}
	return "undefined MessageEntry \n"
}

func (m *Message) Output() string {
	entries := []string{}

	start := fmt.Sprintf(`message %s`, m.Name)

	entries = append(entries, start)

	for _, entry := range m.Entries {
		entries = append(entries, fmt.Sprintf("%s\n", entry.Output()))
	}

	entries = append(entries, "}")
	return strings.Join(entries, "")
}
