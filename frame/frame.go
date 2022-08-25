package frame

import "github.com/danielrjohnson/glVM/values"

type Frame struct {
	retAddr uint64
	locals  map[string]values.Value
}

func New(ra uint64) *Frame {
	return &Frame{retAddr: ra, locals: make(map[string]values.Value)}
}

func (f *Frame) RetAddr() uint64 {
	return f.retAddr
}

func (f *Frame) LocalVars() map[string]values.Value {
	return f.locals
}

func (f *Frame) GetLocal(name string) (values.Value, bool) {
	local, exists := f.locals[name]
	return local, exists
}

func (f *Frame) SetLocal(name string, val values.Value) {
	f.locals[name] = val
}
