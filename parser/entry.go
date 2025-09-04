package parser

import (
	"fmt"
	"strings"

	"github.com/alecthomas/participle/v2/lexer"
)

type Entry struct {
	Pos lexer.Position

	Syntax  string   `  "syntax" "=" @String`
	Package string   `| "package" @(Ident ( "." Ident )*)`
	Import  string   `| "import" @String`
	Message *Message `| @@`
	Service *Service `| @@`
	Enum    *Enum    `| @@`
	Option  *Option  `| "option" @@`
	Extend  *Extend  `| @@`
}

func (e *Entry) Output() string {
	entries := []string{}

	if e.Syntax != "" {
		entries = append(entries, fmt.Sprintf("syntax = %s", e.Syntax))
	}

	if e.Package != "" {
		entries = append(entries, fmt.Sprintf("package %s", e.Package))
	}

	if e.Import != "" {
		entries = append(entries, fmt.Sprintf("import %s", e.Import))
	}

	if e.Message != nil {
		entries = append(entries, e.Message.Output())
	}

	if e.Enum != nil {
		entries = append(entries, e.Enum.Output())
	}

	return strings.Join(entries, "\n")
}
