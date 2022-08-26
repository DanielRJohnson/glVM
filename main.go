package main

import (
	"fmt"
	"os"

	"github.com/danielrjohnson/glVM/VM"
	"github.com/danielrjohnson/glVM/parser"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: glvm <filename>")
		os.Exit(1)
	}
	file, err := os.Open(args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	program := parser.ParseFile(file)
	vm := VM.New(program)
	vm.Run()
	fmt.Println(vm.Show())
}
