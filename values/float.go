package values

import (
	"fmt"

	"github.com/danielrjohnson/glVM/types"
)

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
