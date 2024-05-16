package encoding

import (
    "avr-asm/ast"
    "avr-asm/parser"
)

func Encode(inst ast.Instruction) uint16 {

}

type Operand struct {
    Type       OperandType
    Constraint OperandConstraint
}

type Constraint func(ast.Node) bool

func IsRegister(ast.Node) bool {
    return ast.(type) == ast.Register
}

New(IsRegister, Between(R0, R15))

