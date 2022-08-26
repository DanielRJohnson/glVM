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
	// main
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(3)})
	prog.PushInstruction(instructions.CALL, []values.Value{values.FromString("fib")})
	prog.PushInstruction(instructions.RET, []values.Value{})

	// fib function
	prog.PushLabel("fib")
	prog.PushInstruction(instructions.SET, []values.Value{values.FromString("n")}) // set arg as local
	// fib(0) = 0
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(0)})
	prog.PushInstruction(instructions.GET, []values.Value{values.FromString("n")})
	prog.PushInstruction(instructions.JNE, []values.Value{values.FromString("NotZero")})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(0)})
	prog.PushInstruction(instructions.RET, []values.Value{})
	// fib(1) = 1
	prog.PushLabel("NotZero")
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(1)})
	prog.PushInstruction(instructions.GET, []values.Value{values.FromString("n")})
	prog.PushInstruction(instructions.JNE, []values.Value{values.FromString("NotOne")})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(1)})
	prog.PushInstruction(instructions.RET, []values.Value{})
	// fib(n) = fib(n-1) + fib(n-2)
	prog.PushLabel("NotOne")
	prog.PushInstruction(instructions.GET, []values.Value{values.FromString("n")})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(1)})
	prog.PushInstruction(instructions.SUB, []values.Value{})
	prog.PushInstruction(instructions.CALL, []values.Value{values.FromString("fib")})

	prog.PushInstruction(instructions.GET, []values.Value{values.FromString("n")})
	prog.PushInstruction(instructions.PUSH, []values.Value{values.FromInt(2)})
	prog.PushInstruction(instructions.SUB, []values.Value{})
	prog.PushInstruction(instructions.CALL, []values.Value{values.FromString("fib")})

	prog.PushInstruction(instructions.ADD, []values.Value{})
	prog.PushInstruction(instructions.RET, []values.Value{})

	out := prog.Disassemble()
	fmt.Println(out)

	vm := VM.New(prog)
	vm.Show()
	for vm.IPInRange() {
		vm.Step()
		fmt.Println(vm.Show())
	}
}
