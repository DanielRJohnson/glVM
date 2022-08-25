package values

import (
	"testing"

	"github.com/danielrjohnson/glVM/types"
	"github.com/stretchr/testify/assert"
)

func Test_FromFloatReturnsValueWithCorrectFloat(t *testing.T) {
	val := FromFloat(5.0)
	assert.Equal(t, Float{value: 5.0}, val, "FromFloat did not return correct Value")
}

func Test_ValueReturnsValueFloat(t *testing.T) {
	val := FromFloat(5.0)
	assert.Equal(t, float32(5.0), val.Value().(float32), "Value did not return correct Value")
}

func Test_TypeReturnsTypeFloat(t *testing.T) {
	val := FromFloat(5.0)
	assert.Equal(t, types.Type(types.Float), val.Type(), "Type did not return correct Type")
}

func Test_StringReturnsReprFloat(t *testing.T) {
	val := FromFloat(5.0)
	assert.Equal(t, "5.000000", val.String(), "String did not return correct Repr")
}
