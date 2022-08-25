package values

import "github.com/danielrjohnson/glVM/types"

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
