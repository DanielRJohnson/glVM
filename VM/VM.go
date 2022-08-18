package VM

import (
	"fmt"

	"github.com/danielrjohnson/glVM/VM/stack"
	"github.com/danielrjohnson/glVM/instructions"
	"github.com/danielrjohnson/glVM/program"
	"github.com/danielrjohnson/glVM/values"
)

type VM struct {
	program program.Program
	ip      uint64
	stack   stack.Stack[values.Value]
}

func New(program program.Program) VM {
	return VM{
		program,
		0,
		stack.New[values.Value](),
	}
}

func (vm *VM) Run() {
	code := vm.program.Code()
	data := vm.program.Data()
	for int(vm.ip) < len(code) {
		instr := code[vm.ip]
		switch instr {
		case instructions.NOOP:
			fmt.Println("NOOP :)")
		case instructions.PUSH:
			vm.ip++
			vm.stack.Push(data[code[vm.ip]])
		case instructions.ADD:
			op1, op2, _ := vm.stack.Pop2()
			if op1.Type() == "int" && op2.Type() == "int" {
				vm.stack.Push(values.FromInt(op1.Value().(int) + op2.Value().(int)))
			}
		}
		vm.ip++
	}
}

func (vm VM) Show() {
	fmt.Println("===")
	fmt.Print("CODE | ")
	for i, part := range vm.program.Code() {
		if i == int(vm.ip) {
			fmt.Print("->")
		}
		fmt.Print(part, " | ")
	}
	fmt.Println()
	fmt.Print("DATA | ")
	for _, data := range vm.program.Data() {
		fmt.Print(data, " | ")
	}
	fmt.Println()
	fmt.Print("STACK | ")
	for _, op := range vm.stack.Items() {
		fmt.Print(op, " | ")
	}
	fmt.Println()
	fmt.Println("===")
}
