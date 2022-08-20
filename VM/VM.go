package VM

import (
	"fmt"

	"github.com/danielrjohnson/glVM/VM/stack"
	"github.com/danielrjohnson/glVM/instructions"
	"github.com/danielrjohnson/glVM/program"
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
		instructions.ADD:  vm.Add,
	}
	return &vm
}

func (vm *VM) Run() {
	for int(vm.ip) < len(vm.program.Code()) {
		instr := vm.GetCodeAtIP()
		vm.instructionTable[instr]()
		vm.AdvanceIP()
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

func (vm *VM) Noop() {}

func (vm *VM) Push() {
	vm.AdvanceIP()
	data := vm.GetDataFromIP()
	vm.stack.Push(data)
}

func (vm *VM) Add() {
	op1, op2, _ := vm.stack.Pop2()
	if op1.Type() == "int" && op2.Type() == "int" {
		result := op1.Value().(int) + op2.Value().(int)
		vm.stack.Push(values.FromInt(result))
	}
}

func (vm *VM) Show() {
	fmt.Println("=== VM ===")
	fmt.Print("CODE | ")
	for i, part := range vm.program.Code() {
		if i == int(vm.ip) {
			fmt.Print("->")
		}
		fmt.Print(part, " | ")
	}
	fmt.Print("\nDATA | ")
	for _, data := range vm.program.Data() {
		fmt.Print(data, " | ")
	}
	fmt.Print("\nSTACK | ")
	for _, op := range vm.stack.Items() {
		fmt.Print(op, " | ")
	}
	fmt.Println("\n==========")
}
