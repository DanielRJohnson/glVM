package values

import (
	"fmt"

	"github.com/danielrjohnson/glVM/types"
)

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
