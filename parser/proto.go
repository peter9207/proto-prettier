package parser

import (
	"fmt"
	"io"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Proto struct {
	Pos lexer.Position

	Entries []*Entry `( @@ ";"* )*`
}

type Array struct {
	Pos lexer.Position

	Elements []*Value `"[" ( @@ ( ","? @@ )* )? "]"`
}

type Map struct {
	Pos lexer.Position

	Entries []*MapEntry `"{" ( @@ ( ( "," )? @@ )* )? "}"`
}

type MapEntry struct {
	Pos lexer.Position

	Key   *Value `@@`
	Value *Value `":"? @@`
}

type Extensions struct {
	Pos lexer.Position

	Extensions []Range `"extensions" @@ ( "," @@ )*`
}

type Reserved struct {
	Pos lexer.Position

	Reserved []Range `"reserved" @@ ( "," @@ )*`
}

type Range struct {
	Ident string `  @String`
	Start int    `| ( @Int`
	End   *int   `  ( "to" ( @Int`
	Max   bool   `           | @"max" ) )? )`
}

type Extend struct {
	Pos lexer.Position

	Reference string   `"extend" @Ident ( "." @Ident )*`
	Fields    []*Field `"{" ( @@ ";"? )* "}"`
}

type Service struct {
	Pos lexer.Position

	Name  string          `"service" @Ident`
	Entry []*ServiceEntry `"{" ( @@ ";"? )* "}"`
}

type ServiceEntry struct {
	Pos lexer.Position

	Option *Option `  "option" @@`
	Method *Method `| @@`
}

type Method struct {
	Pos lexer.Position

	Name              string    `"rpc" @Ident`
	StreamingRequest  bool      `"(" @"stream"?`
	Request           *Type     `    @@ ")"`
	StreamingResponse bool      `"returns" "(" @"stream"?`
	Response          *Type     `              @@ ")"`
	Options           []*Option `( "{" ( "option" @@ ";" )* "}" )?`
}

type Oneof struct {
	Pos lexer.Position

	Name    string        `"oneof" @Ident`
	Entries []*OneofEntry `"{" ( @@ ";"* )* "}"`
}

type OneofEntry struct {
	Pos lexer.Position

	Field  *Field  `  @@`
	Option *Option `| "option" @@`
}

func (s Scalar) GoString() string { return scalarToString[s] }

var stringToScalar = map[string]Scalar{
	"double": Double, "float": Float, "int32": Int32, "int64": Int64, "uint32": Uint32, "uint64": Uint64,
	"sint32": Sint32, "sint64": Sint64, "fixed32": Fixed32, "fixed64": Fixed64, "sfixed32": SFixed32,
	"sfixed64": SFixed64, "bool": Bool, "string": String, "bytes": Bytes,
}

func (s *Scalar) Parse(lex *lexer.PeekingLexer) error {
	token := lex.Peek()
	v, ok := stringToScalar[token.Value]
	if !ok {
		return participle.NextMatch
	}
	lex.Next()
	*s = v
	return nil
}

func (p *Proto) Output() string {
	entries := []string{}

	for _, entry := range p.Entries {
		entries = append(entries, fmt.Sprintf("%s\n", entry.Output()))
	}

	return strings.Join(entries, "")
}
func Parse(data io.Reader) (p *Proto, err error) {

	parser := participle.MustBuild[Proto](participle.UseLookahead(2))
	fmt.Println("parser built")

	proto, err := parser.Parse("", data)

	fmt.Println(proto.Output())
	p = proto
	return
}
