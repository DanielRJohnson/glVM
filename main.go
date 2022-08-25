package main

import (
	"fmt"

	"github.com/danielrjohnson/glVM/VM"
	"github.com/danielrjohnson/glVM/instructions"
	"github.com/danielrjohnson/glVM/program"
	"github.com/danielrjohnson/glVM/values"
)

func main() {
	prog := program.New()
	prog.PushInstruction(instructions.NOOP, []values.Value{})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(5)})
	prog.PushInstruction(instructions.ADD, []values.Value{})
	out := prog.Disassemble()
	fmt.Println(out)

	vm := VM.New(prog)
	fmt.Println(vm.Show())
	vm.Run()
	fmt.Println(vm.Show())
}
