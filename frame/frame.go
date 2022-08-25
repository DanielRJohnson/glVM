package frame

type Frame struct {
	retAddr uint64
}

func New(ra uint64) Frame {
	return Frame{retAddr: ra}
}

func (f *Frame) RetAddr() uint64 {
	return f.retAddr
}
