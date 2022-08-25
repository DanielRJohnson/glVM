package values

import (
	"github.com/danielrjohnson/glVM/types"
)

type Value interface {
	Value() any
	Type() types.Type
	String() string
}
