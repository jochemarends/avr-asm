package encoding

import (
    "avr-asm/ast"
    "avr-asm/parser"
)

func Encode(inst ast.Instruction) uint16 {
    opcode := Opcodes[ast.Mnmonic]
    operands := Operands[opcode]
}

type Operand struct {
    Type       OperandType
    Constraint OperandConstraint
    
}

func (Operand) Encode(node ast.Operand) uint16 {
    for _, c := range constraints
}

const (
    ByteRegister OperandType = iota
    WordRegister
)

type Constraint func(ast.Operand) (uint16, error)
type Encoder func(uint16) uint16


func IsRegister(ast.Node) bool {
    return ast.(type) == ast.Register
}

for _, c := range constraints {
    c.Encode(node)
}

func Add()

op, err := op.Encode(node)

// type check -> map operand -> map instruction

New(IsRegister, Between(R0, R15))

type OperandType uint

