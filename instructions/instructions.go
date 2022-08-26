package instructions

const (
	NOOP = iota
	PUSH
	CALL
	RET
	SET
	GET
	ADD
	SUB
	MUL
	DIV
	JE
	JNE
	J
)

var InstructionNames = map[int]string{
	NOOP: "NOOP",
	PUSH: "PUSH",
	CALL: "CALL",
	RET:  "RET",
	SET:  "SET",
	GET:  "GET",
	ADD:  "ADD",
	SUB:  "SUB",
	MUL:  "MUL",
	DIV:  "DIV",
	JE:   "JE",
	JNE:  "JNE",
	J:    "J",
}

var NameToInstruction = map[string]int{
	"NOOP": NOOP,
	"PUSH": PUSH,
	"CALL": CALL,
	"RET":  RET,
	"SET":  SET,
	"GET":  GET,
	"ADD":  ADD,
	"SUB":  SUB,
	"MUL":  MUL,
	"DIV":  DIV,
	"JE":   JE,
	"JNE":  JNE,
	"J":    J,
}

var InstructionArities = map[int]int{
	NOOP: 0,
	PUSH: 1,
	CALL: 1,
	RET:  0,
	SET:  1,
	GET:  1,
	ADD:  0,
	SUB:  0,
	MUL:  0,
	DIV:  0,
	JE:   1,
	JNE:  1,
	J:    1,
}
