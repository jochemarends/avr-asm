package ast

import "avr-asm/token"

type (
    Node interface {

    }
)
type Node interface {
    Token() token.Token
}

type Statement interface {
    Node
    Instruction
}

type Instruction struct {
    Mnemonic Symbol
    Operands []Operand
}

type Operand interface {
    Node
    operandNode()
}

type Register struct {
    token Token
}

type RegisterPair struct {
    Lower Register
    Upper Register
}

// Not a fan of this method of constraining types, but Go seems to 
// be doing it for their "go/ast" package, so there's probably no
// better way.
func (Register)     operandNode() {}
func (RegisterPair) operandNode() {}

type SizeType byte

const (
    Byte SizeType iota 
    Word
)

