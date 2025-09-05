package parser

import (
	"fmt"

	"github.com/alecthomas/participle/v2/lexer"
)

type Type struct {
	Pos lexer.Position

	Scalar    Scalar   `  @@`
	Map       *MapType `| @@`
	Reference string   `| @(Ident ( "." Ident )*)`
}

type MapType struct {
	Pos lexer.Position

	Key   *Type `"map" "<" @@`
	Value *Type `"," @@ ">"`
}

func (t *Type) Output() string {

	if t.Scalar != None {
		return scalarOutput[t.Scalar]
	}

	if t.Map != nil {
		return fmt.Sprintf("map<%s, %s>", t.Map.Key.Output(), t.Map.Value.Output())
	}

	return "unknown Type output\n"
}

var scalarOutput = map[Scalar]string{
	None: "none", Double: "double", Float: "float", Int32: "int32", Int64: "int64", Uint32: "uint32",
	Uint64: "uint64", Sint32: "sint32", Sint64: "sint64", Fixed32: "fixed32", Fixed64: "fixed64",
	SFixed32: "sfixed32", SFixed64: "sfixed64", Bool: "bool", String: "string", Bytes: "bytes",
}

type Scalar int

const (
	None Scalar = iota
	Double
	Float
	Int32
	Int64
	Uint32
	Uint64
	Sint32
	Sint64
	Fixed32
	Fixed64
	SFixed32
	SFixed64
	Bool
	String
	Bytes
)

var scalarToString = map[Scalar]string{
	None: "None", Double: "Double", Float: "Float", Int32: "Int32", Int64: "Int64", Uint32: "Uint32",
	Uint64: "Uint64", Sint32: "Sint32", Sint64: "Sint64", Fixed32: "Fixed32", Fixed64: "Fixed64",
	SFixed32: "SFixed32", SFixed64: "SFixed64", Bool: "Bool", String: "String", Bytes: "Bytes",
}
