package values

import (
	"fmt"

	"github.com/danielrjohnson/glVM/types"
)

type Value interface {
	Value() any
	Type() types.Type
	String() string
}

type Int struct {
	value int
}

func FromInt(i int) Value {
	return Int{value: i}
}

func (i Int) Value() any {
	return i.value
}

func (i Int) Type() types.Type {
	return types.Int
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i.value)
}

type Float struct {
	value float32
}

func FromFloat(f float32) Value {
	return Float{value: f}
}

func (f Float) Value() any {
	return f.value
}

func (f Float) Type() types.Type {
	return types.Float
}

func (f Float) String() string {
	return fmt.Sprintf("%f", f.value)
}

type String struct {
	value string
}

func FromString(s string) Value {
	return String{value: s}
}

func (s String) Value() any {
	return s.value
}

func (s String) Type() types.Type {
	return types.String
}

func (s String) String() string {
	return s.value
}
