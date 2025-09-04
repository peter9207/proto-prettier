package parser

import (
	"github.com/alecthomas/participle"
)

type Proto struct{
	Import Import 
	Messages []Message
	Enums []Enum
}

type Import struct {
	Value string `@@*`
}

type Message struct{
	Fields []FieldDef 
	Reserved []ReservedDef
}


type FieldDef struct {
	Offset int64
	Name string
	Type string
}

type Enum struct {
	Fields []FieldDef
}


type ReservedDef struct{

}

func Parse(data string) (p Proto, err error) {
	parser, err := participle.Build(&p)
	if err != nil {
		return
	}

	err = parser.ParseString(data, &p)
	return
}