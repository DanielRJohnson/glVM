package VM

import (
	"bytes"
	"fmt"

	"github.com/danielrjohnson/glVM/instructions"
	"github.com/danielrjohnson/glVM/program"
	"github.com/danielrjohnson/glVM/stack"
	"github.com/danielrjohnson/glVM/types"
	"github.com/danielrjohnson/glVM/values"
)

type VM struct {
	program          program.Program
	ip               uint64
	stack            stack.Stack[values.Value]
	instructionTable map[int]func()
}

func New(program program.Program) *VM {
	vm := VM{
		program,
		0,
		stack.New[values.Value](),
		make(map[int]func()),
	}
	vm.instructionTable = map[int]func(){
		instructions.NOOP: vm.Noop,
		instructions.PUSH: vm.Push,
		instructions.J:    vm.J,
		instructions.JE:   vm.JE,
		instructions.ADD:  vm.Add,
		instructions.SUB:  vm.Sub,
		instructions.MUL:  vm.Mul,
		instructions.DIV:  vm.Div,
	}
	return &vm
}

func (vm *VM) Run() {
	for int(vm.ip) < len(vm.program.Code()) {
		vm.Step()
	}
}

func (vm *VM) Step() {
	instr := vm.GetCodeAtIP()
	vm.instructionTable[instr]()
	vm.AdvanceIP()
}

func (vm *VM) Noop() {}

func (vm *VM) Push() {
	vm.AdvanceIP()
	data := vm.GetDataFromIP()
	vm.stack.Push(data)
}

func (vm *VM) J() {
	vm.AdvanceIP()
	label := vm.GetDataFromIP()
	vm.ip = uint64(vm.program.Labels()[label.Value().(string)])
}

func (vm *VM) JE() {
	vm.AdvanceIP()
	label := vm.GetDataFromIP()
	op1, op2, _ := vm.stack.Pop2()
	if op1 == op2 {
		vm.ip = uint64(vm.program.Labels()[label.Value().(string)])
	}
}

func (vm *VM) Add() {
	op2, op1, _ := vm.stack.Pop2()
	if op1.Type() == types.Int && op2.Type() == types.Int {
		result := op1.Value().(int) + op2.Value().(int)
		vm.stack.Push(values.FromInt(result))
	} else if op1.Type() == types.Float && op2.Type() == types.Int {
		result := op1.Value().(float32) + float32(op2.Value().(int))
		vm.stack.Push(values.FromFloat(result))
	} else if op1.Type() == types.Int && op2.Type() == types.Float {
		result := float32(op1.Value().(int)) + op2.Value().(float32)
		vm.stack.Push(values.FromFloat(result))
	} else if op1.Type() == types.Float && op2.Type() == types.Float {
		result := op1.Value().(float32) + op2.Value().(float32)
		vm.stack.Push(values.FromFloat(result))
	} else if op1.Type() == types.String && op2.Type() == types.String {
		result := op1.Value().(string) + op2.Value().(string)
		vm.stack.Push(values.FromString(result))
	}
}

func (vm *VM) Sub() {
	op2, op1, _ := vm.stack.Pop2()
	if op1.Type() == types.Int && op2.Type() == types.Int {
		result := op1.Value().(int) - op2.Value().(int)
		vm.stack.Push(values.FromInt(result))
	} else if op1.Type() == types.Float && op2.Type() == types.Int {
		result := op1.Value().(float32) - float32(op2.Value().(int))
		vm.stack.Push(values.FromFloat(result))
	} else if op1.Type() == types.Int && op2.Type() == types.Float {
		result := float32(op1.Value().(int)) - op2.Value().(float32)
		vm.stack.Push(values.FromFloat(result))
	} else if op1.Type() == types.Float && op2.Type() == types.Float {
		result := op1.Value().(float32) - op2.Value().(float32)
		vm.stack.Push(values.FromFloat(result))
	}
}

func (vm *VM) Mul() {
	op2, op1, _ := vm.stack.Pop2()
	if op1.Type() == types.Int && op2.Type() == types.Int {
		result := op1.Value().(int) * op2.Value().(int)
		vm.stack.Push(values.FromInt(result))
	} else if op1.Type() == types.Float && op2.Type() == types.Int {
		result := op1.Value().(float32) * float32(op2.Value().(int))
		vm.stack.Push(values.FromFloat(result))
	} else if op1.Type() == types.Int && op2.Type() == types.Float {
		result := float32(op1.Value().(int)) * op2.Value().(float32)
		vm.stack.Push(values.FromFloat(result))
	} else if op1.Type() == types.Float && op2.Type() == types.Float {
		result := op1.Value().(float32) * op2.Value().(float32)
		vm.stack.Push(values.FromFloat(result))
	}
}

func (vm *VM) Div() {
	op2, op1, _ := vm.stack.Pop2()
	if op1.Type() == types.Int && op2.Type() == types.Int {
		result := op1.Value().(int) / op2.Value().(int)
		vm.stack.Push(values.FromInt(result))
	} else if op1.Type() == types.Float && op2.Type() == types.Int {
		result := op1.Value().(float32) / float32(op2.Value().(int))
		vm.stack.Push(values.FromFloat(result))
	} else if op1.Type() == types.Int && op2.Type() == types.Float {
		result := float32(op1.Value().(int)) / op2.Value().(float32)
		vm.stack.Push(values.FromFloat(result))
	} else if op1.Type() == types.Float && op2.Type() == types.Float {
		result := op1.Value().(float32) / op2.Value().(float32)
		vm.stack.Push(values.FromFloat(result))
	}
}

func (vm *VM) GetCodeAtIP() int {
	return vm.program.Code()[vm.ip]
}

func (vm *VM) GetDataFromIP() values.Value {
	dataIdx := vm.GetCodeAtIP()
	return vm.program.Data()[dataIdx]
}

func (vm *VM) AdvanceIP() {
	vm.ip++
}

func (vm *VM) Show() string {
	var buf bytes.Buffer
	buf.WriteString("=== VM ===\n")
	buf.WriteString("CODE | ")
	for i, part := range vm.program.Code() {
		if i == int(vm.ip) {
			buf.WriteString("->")
		}
		buf.WriteString(fmt.Sprintf("%d | ", part))
	}
	buf.WriteString("\nDATA | ")
	for _, data := range vm.program.Data() {
		buf.WriteString(fmt.Sprintf("%d | ", data))
	}
	buf.WriteString("\nSTACK | ")
	for _, op := range vm.stack.Items() {
		buf.WriteString(fmt.Sprintf("%d | ", op))
	}
	buf.WriteString("\n==========\n")
	return buf.String()
}
