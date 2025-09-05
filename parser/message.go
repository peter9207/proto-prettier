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

func (m *Message) Output() string {
	entries := []string{}

	start := fmt.Sprintf("message %s {\n", m.Name)

	entries = append(entries, start)

	for _, entry := range m.Entries {
		entries = append(entries, fmt.Sprintf("%s;\n", entry.Output()))
	}

	entries = append(entries, "}")
	return strings.Join(entries, "")
}
