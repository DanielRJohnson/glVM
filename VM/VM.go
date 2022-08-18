package VM

import (
	"fmt"

	"github.com/danielrjohnson/glVM/VM/stack"
	"github.com/danielrjohnson/glVM/instructions"
	"github.com/danielrjohnson/glVM/program"
)

type VM[T any] struct {
	program program.Program[T]
	ip      uint64
	stack   stack.Stack[T]
}

func New[T any](program program.Program[T]) VM[T] {
	return VM[T]{
		program,
		0,
		stack.New[T](),
	}
}

func (vm VM[T]) Run() {
	code := vm.program.Code()
	for int(vm.ip) < len(code) {
		instr := code[vm.ip]
		switch instr {
		case instructions.NOOP:
			fmt.Println("NOOP :)")
		}
		vm.ip++
	}
}
