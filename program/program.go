package program

import (
	"bytes"
	"fmt"

	"github.com/danielrjohnson/glVM/instructions"
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

func (p *Program) PushInstruction(opcode int, args []values.Value) error {
	p.code = append(p.code, opcode)
	for _, arg := range args {
		dataIdx := p.PushData(arg)
		p.code = append(p.code, dataIdx)
	}
	return nil
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
		out.WriteString(fmt.Sprintf("@%d: %v\n", idx, data))
	}
	out.WriteString("\n")

	reversedLabels := make(map[int]string)
	for name, ip := range p.labels {
		reversedLabels[ip] = name
	}

	for i := 0; i < len(p.code); i++ {
		if name, ok := reversedLabels[i]; ok {
			out.WriteString(fmt.Sprintf("%s: \n", name))
		}
		out.WriteString(instructions.InstructionNames[p.code[i]])
		for j := 0; j < instructions.InstructionArities[p.code[i]]; j++ {
			out.WriteString(fmt.Sprintf(" @%d ", p.code[i]))
			i++
		}
		out.WriteString("\n")
	}
	return out.String()
}
