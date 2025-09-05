package parser

import (
	"fmt"
	"strings"

	"github.com/alecthomas/participle/v2/lexer"
)

type Field struct {
	Pos lexer.Position

	Optional bool `(   @"optional"`
	Required bool `  | @"required"`
	Repeated bool `  | @"repeated" )?`

	Type *Type  `@@`
	Name string `@Ident`
	Tag  int    `"=" @Int`

	Options []*Option `( "[" @@ ( "," @@ )* "]" )?`
}

func (f *Field) Output() string {

	entries := []string{}

	if f.Optional {
		entries = append(entries, "optional")
	}

	if f.Required {
		entries = append(entries, "required")
	}

	if f.Repeated {
		entries = append(entries, "repeated")
	}

	entries = append(entries, fmt.Sprintf("%s %s = %d", f.Type.Output(), f.Name, f.Tag))

	return strings.Join(entries, " ")
}
