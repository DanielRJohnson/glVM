# glVM

glVM (short for Glimmer VM) is a stack-based virtual machine with a concise instruction set, bytecode file parsing,
and an exposed API for building programs at compile time. The bytecode that this machine runs has support for 
function calls, control flow, and int/float/string computations. The machine itself is also tested at 100% code coverage.

## Instruction Set

| Instruction | Arity | Description                                                                           |
| ----------- | ----- | ------------------------------------------------------------------------------------- |
| NOOP        | 0     | Do nothing                                                                            |
| PUSH        | 1     | Push a given value onto the stack                                                     |
| CALL        | 1     | Push a new frame on the call stack and jump to the given label                        |
| RET         | 0     | Pop a frame off of the call stack and jump to its return address                      |
| SET         | 1     | Pops the stack and binds that value to a given name in the current frame              |
| GET         | 1     | Push the value of a given name in the current frame onto the stack                    |
| ADD         | 0     | Pop 2 values off of the stack and push their sum                                      |
| SUB         | 0     | Pop 2 values off of the stack and push their difference                               |
| MUL         | 0     | Pop 2 values off of the stack and push their product                                  |
| DIV         | 0     | Pop 2 values off of the stack and push their quotient                                 |
| JE          | 1     | Pop 2 values off of the stack and jump to the given label if the values are equal     |
| JNE         | 1     | Pop 2 values off of the stack and jump to the given label if the values are not equal |
| J           | 1     | Unconditionally jump to the given label                                               |

## Example Program, Fibbonacci

The following code snippet is a program that calculates the 20th fibbonacci number using recursive calls. You can try this yourself by having an installation of Go 1.18, cloning this repo, and running `./glVM _examples/fib.gbc`. Try changing the number, but note that the spacing must be exactly like the given example as the parser is very simple.

```
@0: I(20)
@1: S("fib")
@2: S("n")
@3: I(0)
@4: S("NotZero")
@5: I(1)
@6: S("NotOne")
@7: I(2)

#main: PUSH @0
CALL @1
RET
#fib: SET @2
PUSH @3
GET @2
JNE @4
PUSH @3
RET
#NotZero: PUSH @5
GET @2
JNE @6
PUSH @5
RET
#NotOne: GET @2
PUSH @5
SUB
CALL @1
GET @2
PUSH @7
SUB
CALL @1
ADD
RET
```

## TODO

~~Functioning machine that executes instructions, has stored data & labels~~

~~Push instruction to add item to stack~~

~~Integers + integer arithmetic~~

~~Floats + float arithmetic~~

~~Type promotion for int+float arithmetic~~

~~Strings + string concatenation~~

~~Conditionals + loops (jump structure of some sort, JE)~~

~~functions (call, jump, ret, call stack, fib test)~~

~~/program tests~~

~~/values tests~~

~~/frame tests~~

~~Disassembly for program~~

~~Bytecode file reading/execution (main executable)~~

~~Pimp out readme (instruction table, fib program, etc)~~

## Not Planned Currently, but Reasonable Extensions

Booleans + boolean logic

Arrays + array operations

Standard Library (IO, helpers, etc)

Some sort of structs or classes

Make a compiler from a language to this bytecode

etc
