package frame

import (
	"testing"

	"github.com/danielrjohnson/glVM/values"
	"github.com/stretchr/testify/assert"
)

func Test_NewReturnsEmptyFrame(t *testing.T) {
	frame := New(2)
	assert.Equal(t, uint64(2), frame.retAddr, "New frame has incorrect retAddr")
	assert.Equal(t, make(map[string]values.Value), frame.locals, "New frame has non-empty locals")
}

func Test_RetAddrReturnsRetAddr(t *testing.T) {
	frame := New(2)
	assert.Equal(t, uint64(2), frame.RetAddr(), "RetAddr returned incorrect retAddr")
}

func Test_GetLocalGetsLocalFromLocals(t *testing.T) {
	frame := New(2)
	frame.SetLocal("Hello", values.FromString("World"))
	local, _ := frame.GetLocal("Hello")
	assert.Equal(t, values.FromString("World"), local, "SetLocal did not set local in locals")
}

func Test_SetLocalSetsLocalInLocals(t *testing.T) {
	frame := New(2)
	frame.SetLocal("Hello", values.FromString("World"))
	assert.Equal(t, values.FromString("World"), frame.locals["Hello"], "SetLocal did not set local in locals")
}
