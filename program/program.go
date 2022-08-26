package program

import (
	"bytes"
	"fmt"

	"github.com/danielrjohnson/glVM/instructions"
	"github.com/danielrjohnson/glVM/types"
	"github.com/danielrjohnson/glVM/values"
)

type Program struct {
	code   []int
	data   []values.Value
	labels map[string]int
}

func New() Program {
	program := Program{
		code:   []int{},
		data:   []values.Value{},
		labels: make(map[string]int),
	}
	program.labels["main"] = 0
	return program
}

func (p *Program) PushInstruction(opcode int, args []values.Value) {
	p.code = append(p.code, opcode)
	for _, arg := range args {
		dataIdx := p.PushData(arg)
		p.code = append(p.code, dataIdx)
	}
}

func (p *Program) PushInstructionRawDataIdx(opcode int, args []int) {
	p.code = append(p.code, opcode)
	p.code = append(p.code, args...)
}

func (p *Program) PushData(item values.Value) int {
	p.data = append(p.data, item)
	return len(p.data) - 1
}

func (p *Program) PushLabel(name string) {
	p.labels[name] = len(p.code)
}

func (p Program) Code() []int {
	return p.code
}

func (p Program) Data() []values.Value {
	return p.data
}

func (p Program) Labels() map[string]int {
	return p.labels
}

func (p Program) Disassemble() string {
	var out bytes.Buffer
	for idx, data := range p.data {
		if data.Type() == types.String {
			out.WriteString(fmt.Sprintf("@%d: S(%q)\n", idx, data))
		} else if data.Type() == types.Int {
			out.WriteString(fmt.Sprintf("@%d: I(%v)\n", idx, data))
		} else if data.Type() == types.Float {
			out.WriteString(fmt.Sprintf("@%d: F(%v)\n", idx, data))
		}
	}
	out.WriteString("\n")

	reversedLabels := make(map[int]string)
	for name, ip := range p.labels {
		reversedLabels[ip] = name
	}

	for i := 0; i < len(p.code); i++ {
		if name, ok := reversedLabels[i]; ok {
			out.WriteString(fmt.Sprintf("#%s: ", name))
		}
		out.WriteString(instructions.InstructionNames[p.code[i]])
		for j := 0; j < instructions.InstructionArities[p.code[i]]; j++ {
			i++
			out.WriteString(fmt.Sprintf(" @%d ", p.code[i+j]))
		}
		out.WriteString("\n")
	}
	return out.String()
}
