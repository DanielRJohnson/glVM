package VM

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danielrjohnson/glVM/instructions"
	"github.com/danielrjohnson/glVM/program"
	"github.com/danielrjohnson/glVM/values"
)

func Test_NewVMHasEmptyStack(t *testing.T) {
	prog := program.New()
	vm := New(prog)
	assert.Truef(t, vm.stack.IsEmpty(), "New VM does not have empty stack, got size %d", vm.stack.Size())
}

func Test_RunExecutesAllInstructions(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	vm := New(prog)
	vm.Run()
	assert.NotEqualf(t, vm.ip, 3, "Run did not execute all instructions resulting in wrong ip, got=%d expected=%d", vm.ip, 3)
}

func Test_StepExecutesOneInstruction(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	vm := New(prog)
	vm.Step()
	assert.NotEqualf(t, vm.ip, 1, "Step did not execute one instruction resulting in wrong ip, got=%d expected=%d", vm.ip, 1)
}

func Test_JeJumpsToLabelIfEqual(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.JE, []values.Value{values.FromString("After")})
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	prog.PushLabel("After")
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	vm := New(prog)
	vm.Step()
	vm.Step()
	vm.Step()
	assert.Equalf(t, vm.ip, uint64(7), "JE did not jump to correct ip when values are equal, got=%d expected=%d", vm.ip, 7)
}

func Test_JeDoesNotJumpToLabelIfNotEqual(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(4)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.JE, []values.Value{values.FromString("After")})
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	prog.PushLabel("After")
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	vm := New(prog)
	vm.Step()
	vm.Step()
	vm.Step()
	assert.Equalf(t, vm.ip, uint64(6), "JE jumped to ip when values were not equal, got=%d expected=%d", vm.ip, 6)
}

func Test_JJumpsToLabel(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.J, []values.Value{values.FromString("After")})
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	prog.PushLabel("After")
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	vm := New(prog)
	vm.Step()
	assert.Equalf(t, vm.ip, uint64(3), "J did not jump to correct ip, got=%d expected=%d", vm.ip, 3)
}

func Test_CallJumpsToLabelAndPushesFrameOnCallStack(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.CALL, []values.Value{values.FromString("Function")})
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	prog.PushLabel("Function")
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	vm := New(prog)
	vm.Step()
	assert.Equalf(t, vm.ip, uint64(3), "Call did not jump to correct ip, got=%d expected=%d", vm.ip, 3)
	assert.Equalf(t, vm.callStack.Size(), 2,
		"Call did not push frame on stack, size got=%d expected=%d", vm.callStack.Size(), 2)
}

func Test_RetJumpsToRetAddrAndPopsFrameOffCallStack(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.CALL, []values.Value{values.FromString("Function")})
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	prog.PushLabel("Function")
	prog.PushInstruction(instructions.RET, []values.Value{})
	vm := New(prog)
	vm.Step()
	vm.Step()
	assert.Equalf(t, vm.ip, uint64(2), "Ret did not jump to correct ip, got=%d expected=%d", vm.ip, 2)
	assert.Equalf(t, vm.callStack.Size(), 1,
		"Ret did not pop frame off stack, size got=%d expected=%d", vm.callStack.Size(), 1)
}

func Test_AddPopsTopTwoIntegersAndPushesThierSumAsInteger(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(6)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(7)})
	prog.PushInstruction(instructions.ADD, []values.Value{})
	vm := New(prog)
	vm.Run()
	if assert.Equalf(t, vm.stack.Size(), 2,
		"Add did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 2) {
		sum := vm.stack.Items()[1].Value().(int)
		assert.Equalf(t, sum, 13, "Add did not result in correct sum, got=%d expected=%d", sum, 13)
	}
}

func Test_AddPopsTopTwoFloatsAndPushesThierSumAsFloat(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromFloat(1.2)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromFloat(1.3)})
	prog.PushInstruction(instructions.ADD, []values.Value{})
	vm := New(prog)
	vm.Run()
	if assert.Equalf(t, vm.stack.Size(), 2,
		"Add did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 2) {
		sum := vm.stack.Items()[1].Value().(float32)
		assert.Equalf(t, sum, float32(2.5), "Add did not result in correct sum, got=%f expected=%f", sum, float32(2.5))
	}
}

func Test_AddPopsTopTwoMixedNumbersAndPushesThierSumAsFloat(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(1)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromFloat(1.3)})
	prog.PushInstruction(instructions.ADD, []values.Value{})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(3)})
	prog.PushInstruction(instructions.ADD, []values.Value{})
	vm := New(prog)
	vm.Run()
	if assert.Equalf(t, vm.stack.Size(), 2,
		"Add did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 2) {
		sum := vm.stack.Items()[1].Value().(float32)
		assert.Equalf(t, sum, float32(5.3), "Add did not result in correct sum, got=%f expected=%f", sum, float32(5.3))
	}
}

func Test_AddPopsTopTwoStringsAndPushesThierConcatenatedString(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromString("Hello ")})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromString("World!")})
	prog.PushInstruction(instructions.ADD, []values.Value{})
	vm := New(prog)
	vm.Run()
	if assert.Equalf(t, vm.stack.Size(), 2,
		"Add did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 2) {
		sum := vm.stack.Items()[1].Value().(string)
		assert.Equalf(t, sum, "Hello World!", "Add did not result in correct sum, got=%f expected=%f", sum, "Hello World!")
	}
}

func Test_SubPopsTopTwoAndPushesTheirDifference(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(4)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(2)})
	prog.PushInstruction(instructions.SUB, []values.Value{})
	vm := New(prog)
	vm.Run()
	if assert.Equalf(t, vm.stack.Size(), 2,
		"Sub did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 2) {
		diff := vm.stack.Items()[1].Value().(int)
		assert.Equalf(t, diff, 2, "Sub did not result in correct difference, got=%d expected=%d", diff, 2)
	}
}

func Test_SubPopsTopTwoFloatsAndPushesThierDifferenceAsFloat(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromFloat(1.3)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromFloat(1.2)})
	prog.PushInstruction(instructions.SUB, []values.Value{})
	vm := New(prog)
	vm.Run()
	if assert.Equalf(t, vm.stack.Size(), 2,
		"Sub did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 2) {
		diff := vm.stack.Items()[1].Value().(float32)
		assert.InEpsilonf(t, diff, float32(0.1), 1e-6, "Sub did not result in correct difference, got=%f expected=%f", diff, float32(0.1))
	}
}

func Test_SubPopsTopTwoMixedNumbersAndPushesThierDifferenceAsFloat(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(2)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromFloat(0.7)})
	prog.PushInstruction(instructions.SUB, []values.Value{})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(1)})
	prog.PushInstruction(instructions.SUB, []values.Value{})
	vm := New(prog)
	vm.Run()
	if assert.Equalf(t, vm.stack.Size(), 2,
		"Sub did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 2) {
		diff := vm.stack.Items()[1].Value().(float32)
		assert.InEpsilonf(t, diff, float32(0.3), 1e-6, "Sub did not result in correct difference, got=%f expected=%f", diff, float32(0.3))
	}
}

func Test_MulPopsTopTwoAndPushesTheirProduct(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(4)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(6)})
	prog.PushInstruction(instructions.MUL, []values.Value{})
	vm := New(prog)
	vm.Run()
	if assert.Equalf(t, vm.stack.Size(), 2,
		"Mul did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 2) {
		product := vm.stack.Items()[1].Value().(int)
		assert.Equalf(t, product, 24, "Mul did not result in correct product, got=%d expected=%d", product, 24)
	}
}

func Test_MulPopsTopTwoFloatsAndPushesTheirProductAsFloat(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromFloat(2.5)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromFloat(4)})
	prog.PushInstruction(instructions.MUL, []values.Value{})
	vm := New(prog)
	vm.Run()
	if assert.Equalf(t, vm.stack.Size(), 2,
		"Mul did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 2) {
		product := vm.stack.Items()[1].Value().(float32)
		assert.Equalf(t, product, float32(10), "Mul did not result in correct sum, got=%f expected=%f", product, float32(10))
	}
}

func Test_MulPopsTopTwoMixedNumbersAndPushesThierProductAsFloat(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(2)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromFloat(2.5)})
	prog.PushInstruction(instructions.MUL, []values.Value{})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(2)})
	prog.PushInstruction(instructions.MUL, []values.Value{})
	vm := New(prog)
	vm.Run()
	if assert.Equalf(t, vm.stack.Size(), 2,
		"Mul did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 2) {
		product := vm.stack.Items()[1].Value().(float32)
		assert.Equalf(t, product, float32(10), "Mul did not result in correct product, got=%f expected=%f", product, float32(10))
	}
}

func Test_DivPopsTopTwoAndPushesTheirQuotient(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(8)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(2)})
	prog.PushInstruction(instructions.DIV, []values.Value{})
	vm := New(prog)
	vm.Run()
	if assert.Equalf(t, vm.stack.Size(), 2,
		"Div did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 2) {
		quotient := vm.stack.Items()[1].Value().(int)
		assert.Equalf(t, quotient, 4, "Div did not result in correct quotient, got=%d expected=%d", quotient, 4)
	}
}

func Test_DivPopsTopTwoFloatsAndPushesThierQuotientAsFloat(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromFloat(2.4)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromFloat(2.0)})
	prog.PushInstruction(instructions.DIV, []values.Value{})
	vm := New(prog)
	vm.Run()
	if assert.Equalf(t, vm.stack.Size(), 2,
		"Div did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 2) {
		quotient := vm.stack.Items()[1].Value().(float32)
		assert.Equalf(t, quotient, float32(1.2), "Div did not result in correct quotient, got=%f expected=%f", quotient, float32(1.2))
	}
}

func Test_DivPopsTopTwoMixedNumbersAndPushesThierQuotientAsFloat(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(1)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromFloat(0.5)})
	prog.PushInstruction(instructions.DIV, []values.Value{})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(2)})
	prog.PushInstruction(instructions.DIV, []values.Value{})
	vm := New(prog)
	vm.Run()
	if assert.Equalf(t, vm.stack.Size(), 2,
		"Div did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 2) {
		quotient := vm.stack.Items()[1].Value().(float32)
		assert.Equalf(t, quotient, float32(1), "Div did not result in correct quotient, got=%f expected=%f", quotient, float32(1))
	}
}

func Test_GetCodeAtIPReturnsCorrectCode(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.ADD, []values.Value{})
	vm := New(prog)
	vm.AdvanceIP()
	assert.Equalf(t, vm.GetCodeAtIP(), instructions.PUSH,
		"GetCodeAtIP returned instruction from wrong index, got=%d expected=%d", vm.GetCodeAtIP(), instructions.PUSH)
}

func Test_GetDataFromIPReturnsCorrectData(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(0)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(1)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(2)})
	vm := New(prog)
	vm.AdvanceIP()
	vm.AdvanceIP() // go past push data
	assert.Equalf(t, vm.GetDataFromIP().Value().(int), 1,
		"GetDataFromIP returned data from wrong index, got=%d expected=%d", vm.GetDataFromIP(), 1)
}

func Test_AdvanceIPAdvancedIP(t *testing.T) {
	prog := program.New()
	vm := New(prog)
	vm.AdvanceIP()
	assert.Equalf(t, int(vm.ip), 1, "AdvanceIP did not increment correctly, one advance yielded %d", vm.ip)
}

func Test_ShowDoesNotCrashLmao(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(0)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(1)})
	prog.PushInstruction(instructions.ADD, []values.Value{})
	vm := New(prog)
	vm.Show()
	vm.Run()
	vm.Show()
}
