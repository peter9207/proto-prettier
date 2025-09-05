package parser

import (
	"fmt"
	"strings"

	// "io"

	// "github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Enum struct {
	Pos lexer.Position

	Name   string       `"enum" @Ident`
	Values []*EnumEntry `"{" ( @@ ( ";" )* )* "}"`
}

func (e *Enum) Output() string {

	entries := []string{}

	start := fmt.Sprintf("enum %s {\n", e.Name)
	entries = append(entries, start)

	for _, value := range e.Values {
		entries = append(entries, fmt.Sprintf("%s;\n", value.Output()))
	}

	entries = append(entries, "}\n")
	return strings.Join(entries, "")
}
