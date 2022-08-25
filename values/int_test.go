package values

import (
	"testing"

	"github.com/danielrjohnson/glVM/types"

	"github.com/stretchr/testify/assert"
)

func Test_FromIntReturnsValueWithCorrectInt(t *testing.T) {
	val := FromInt(5)
	assert.Equal(t, Int{value: 5}, val, "FromInt did not return correct Value")
}

func Test_ValueReturnsValueInt(t *testing.T) {
	val := FromInt(5)
	assert.Equal(t, 5, val.Value().(int), "Value did not return correct Value")
}

func Test_TypeReturnsTypeInt(t *testing.T) {
	val := FromInt(5)
	assert.Equal(t, types.Type(types.Int), val.Type(), "Type did not return correct Type")
}

func Test_StringReturnsReprInt(t *testing.T) {
	val := FromInt(5)
	assert.Equal(t, "5", val.String(), "String did not return correct Repr")
}
