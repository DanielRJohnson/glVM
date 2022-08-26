package program

import (
	"testing"

	"github.com/danielrjohnson/glVM/instructions"
	"github.com/danielrjohnson/glVM/values"
	"github.com/stretchr/testify/assert"
)

func Test_NewCreatesEmptyProgramWithMainAtIP0(t *testing.T) {
	prog := New()
	assert.Equal(t, 0, len(prog.code), "code should be initially empty")
	assert.Equal(t, 0, len(prog.data), "data should be initally empty")
	for k, v := range prog.labels {
		assert.Equal(t, "main", k, "labels should only contain key main")
		assert.Equal(t, 0, v, "labels should only contain value 0")
	}
}

func Test_PushInstructionPushesANoArgumentInstruction(t *testing.T) {
	prog := New()
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	assert.Equal(t, 1, len(prog.code), "PushInstruction did not result in correct code size")
}

func Test_PushInstructionPushesInstructionWithArgumentToData(t *testing.T) {
	prog := New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	assert.Equal(t, 2, len(prog.code), "PushInstruction did not result in correct code size")
	assert.Equal(t, 1, len(prog.data), "PushInstruction did not result in correct data size")
}

func Test_PushInstructionRawDataIdxPushesInstructionsAndIdxs(t *testing.T) {
	prog := New()
	prog.PushInstructionRawDataIdx(instructions.PUSH, []int{2, 3})
	assert.Equal(t, 3, len(prog.code), "PushInstructionRawDataIdx did not result in correct code size")
}

func Test_PushDataPushesDataAndReturnsDataIndex(t *testing.T) {
	prog := New()
	idx := prog.PushData(values.FromInt(5))
	assert.Equalf(t, 1, len(prog.data), "PushData did not result in correct data size")
	assert.Equal(t, 0, idx, "PushData did not return correct data index")
}

func Test_PushLabelAddsNameAndIPToLabelsMap(t *testing.T) {
	prog := New()
	prog.PushLabel("MyLabel")
	assert.Equal(t, 0, prog.labels["MyLabel"], "PushLabel did not add correct ip to labels")
}

func Test_CodeReturnsCode(t *testing.T) {
	prog := New()
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	assert.Equal(t, []int{0}, prog.Code(), "Code did not return the correct code")
}

func Test_DataReturnsData(t *testing.T) {
	prog := New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	assert.Equal(t, []values.Value{values.FromInt(5)}, prog.Data(), "Data did not return the correct data")
}

func Test_LabelsReturnsLabels(t *testing.T) {
	prog := New()
	for k, v := range prog.Labels() {
		assert.Equalf(t, k, "main", "labels contains wrong key")
		assert.Equalf(t, v, 0, "labels contains wrong value")
	}
}

func Test_DissassembleDoesntCrashLmao(t *testing.T) {
	prog := New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromString("Hi")})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromFloat(3.4)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(0)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(1)})
	prog.PushInstruction(instructions.ADD, []values.Value{})
	prog.Disassemble()
}
