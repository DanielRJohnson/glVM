package main

import (
	"github.com/danielrjohnson/glVM/VM"
	"github.com/danielrjohnson/glVM/instructions"
	"github.com/danielrjohnson/glVM/program"
)

func main() {
	prog := program.New[int]()
	prog.PushInstruction(instructions.NOOP, []int{})

	vm := VM.New(prog)
	vm.Run()
}
