package values

import "fmt"

type Value interface {
	Value() any
	Type() string
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

func (i Int) Type() string {
	return "int"
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i.value)
}
