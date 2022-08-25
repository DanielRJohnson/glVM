package values

import (
	"testing"

	"github.com/danielrjohnson/glVM/types"
	"github.com/stretchr/testify/assert"
)

func Test_FromStringReturnsValueWithCorrectString(t *testing.T) {
	val := FromString("hi")
	assert.Equal(t, String{value: "hi"}, val, "FromString did not return correct Value")
}

func Test_ValueReturnsValueString(t *testing.T) {
	val := FromString("hi")
	assert.Equal(t, "hi", val.Value().(string), "Value did not return correct Value")
}

func Test_TypeReturnsTypeString(t *testing.T) {
	val := FromString("hi")
	assert.Equal(t, types.Type(types.String), val.Type(), "Type did not return correct Type")
}

func Test_StringReturnsReprString(t *testing.T) {
	val := FromString("hi")
	assert.Equal(t, "hi", val.String(), "String did not return correct Repr")
}
