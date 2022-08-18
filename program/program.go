package program

type Program[T any] struct {
	code   []int
	data   []T
	labels map[string]int
}

func New[T any]() Program[T] {
	program := Program[T]{
		code:   []int{},
		data:   []T{},
		labels: make(map[string]int),
	}
	program.labels["main"] = 0
	return program
}

func (p *Program[T]) PushInstruction(opcode int, args []T) error {
	p.code = append(p.code, opcode)
	for _, arg := range args {
		dataIdx := p.PushData(arg)
		p.code = append(p.code, dataIdx)
	}
	return nil
}

func (p *Program[T]) PushData(item T) int {
	p.data = append(p.data, item)
	return len(p.data) - 1
}

func (p *Program[T]) PushLabel(name string) {
	p.labels[name] = len(p.code)
}

func (p Program[T]) Code() []int {
	return p.code
}

func (p Program[T]) Data() []T {
	return p.data
}

func (p Program[T]) Labels() map[string]int {
	return p.labels
}
