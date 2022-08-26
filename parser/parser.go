package parser

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/danielrjohnson/glVM/instructions"
	"github.com/danielrjohnson/glVM/program"
	"github.com/danielrjohnson/glVM/values"
)

func ParseFile(file *os.File) program.Program {
	scanner := bufio.NewScanner(file)
	prog := program.New()

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue // empty line
		}
		if line[0] == '@' {
			lineParts := strings.Split(line, " ") // ["@X:", data]
			parseData(lineParts[1], &prog)
		} else {
			if line[0] == '#' {
				line = parseAndTrimLabel(line, &prog)
			}
			lineParts := strings.Split(line, " ")
			instr := instructions.NameToInstruction[lineParts[0]]
			args := parseArgs(lineParts[1:])

			prog.PushInstructionRawDataIdx(instr, args)
		}
	}
	return prog
}

func parseData(data string, prog *program.Program) {
	valueStart := strings.Index(data, "(")
	valueEnd := strings.Index(data, ")")
	if data[0] == 'I' {
		givenInt, _ := strconv.ParseInt(data[valueStart+1:valueEnd], 10, 0)
		prog.PushData(values.FromInt(int(givenInt)))
	} else if data[0] == 'F' {
		givenFloat, _ := strconv.ParseFloat(data[valueStart+1:valueEnd], 32)
		prog.PushData(values.FromFloat(float32(givenFloat)))
	} else if data[0] == 'S' {
		givenString := data[valueStart+2 : valueEnd-1] // cut off quotes
		prog.PushData(values.FromString(givenString))
	}
}

func parseAndTrimLabel(line string, prog *program.Program) string {
	colonIdx := strings.Index(line, ":")
	prog.PushLabel(line[1:colonIdx])
	return line[colonIdx+2:] // cut off label, colon, and space
}

func parseArgs(lineParts []string) []int {
	argsInts := []int{}
	for _, argStr := range lineParts {
		argStr = argStr[1:] // cut off "@"
		num, _ := strconv.ParseInt(argStr, 10, 0)
		argsInts = append(argsInts, int(num))
	}
	return argsInts
}
