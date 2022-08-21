package VM

import (
	"testing"

	"github.com/danielrjohnson/glVM/instructions"
	"github.com/danielrjohnson/glVM/program"
	"github.com/danielrjohnson/glVM/values"
)

func Test_NewVMHasEmptyStack(t *testing.T) {
	prog := program.New()
	vm := New(prog)
	if !vm.stack.IsEmpty() {
		t.Errorf("New VM does not have empty stack, got size %d", vm.stack.Size())
	}
}

func Test_RunExecutesAllInstructions(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	vm := New(prog)
	vm.Run()
	if vm.ip != 3 {
		t.Errorf("Run did not execute all instructions resulting in wrong ip, got=%d expected=%d", vm.ip, 3)
	}
}

func Test_StepExecutesOneInstruction(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	vm := New(prog)
	vm.Step()
	if vm.ip != 1 {
		t.Errorf("Step did not execute one instruction resulting in wrong ip, got=%d expected=%d", vm.ip, 1)
	}
}

func Test_PushAddsItemToStack(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	vm := New(prog)
	vm.Step()
	if vm.stack.Size() != 1 {
		t.Fatalf("Push did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 1)
	}
	stackItem := vm.stack.Items()[0].Value().(int)
	if stackItem != 5 {
		t.Errorf("Push did not push correct value to stack, got=%d expected=%d", stackItem, 5)
	}
}

func Test_AddPopsTopTwoAndPushesTheirSum(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(6)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(7)})
	prog.PushInstruction(instructions.ADD, []values.Value{})
	vm := New(prog)
	vm.Run()
	if vm.stack.Size() != 2 {
		t.Fatalf("Add did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 2)
	}
	sum := vm.stack.Items()[1].Value().(int)
	if sum != 13 {
		t.Fatalf("Add did not result in correct sum, got=%d expected=%d", sum, 13)
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
	if vm.stack.Size() != 2 {
		t.Fatalf("Sub did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 2)
	}
	diff := vm.stack.Items()[1].Value().(int)
	if diff != 2 {
		t.Fatalf("Sub did not result in correct difference, got=%d expected=%d", diff, 2)
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
	if vm.stack.Size() != 2 {
		t.Fatalf("Mul did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 2)
	}
	product := vm.stack.Items()[1].Value().(int)
	if product != 24 {
		t.Fatalf("Add did not result in correct product, got=%d expected=%d", product, 24)
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
	if vm.stack.Size() != 2 {
		t.Fatalf("Div did not result in correct stack size, got=%d expected=%d", vm.stack.Size(), 2)
	}
	quotient := vm.stack.Items()[1].Value().(int)
	if quotient != 4 {
		t.Fatalf("Div did not result in correct quotient, got=%d expected=%d", quotient, 4)
	}
}

func Test_GetCodeAtIPReturnsCorrectCode(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.ADD, []values.Value{})
	vm := New(prog)
	vm.AdvanceIP()
	if vm.GetCodeAtIP() != instructions.PUSH {
		t.Errorf("GetCodeAtIP returned instruction from wrong index, got=%d expected=%d",
			vm.GetCodeAtIP(), instructions.PUSH)
	}
}

func Test_GetDataFromIPReturnsCorrectData(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(0)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(1)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(2)})
	vm := New(prog)
	vm.AdvanceIP()
	vm.AdvanceIP() // go past push data
	if vm.GetDataFromIP().Value().(int) != 1 {
		t.Errorf("GetDataFromIP returned data from wrong index, got=%d expected=%d",
			vm.GetDataFromIP(), 1)
	}
}

func Test_AdvanceIPAdvancedIP(t *testing.T) {
	prog := program.New()
	vm := New(prog)
	vm.AdvanceIP()
	if vm.ip != 1 {
		t.Errorf("AdvanceIP did not increment correctly, one advance yielded %d", vm.ip)
	}
}

func Test_ShowDoesNotCrash(t *testing.T) {
	prog := program.New()
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(0)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(1)})
	prog.PushInstruction(instructions.ADD, []values.Value{})
	vm := New(prog)
	vm.Show()
	vm.Run()
	vm.Show()
}
