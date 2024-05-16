package ast

import "avr-asm/token"

type Node interface {
    //Token() token.Token
}

type Statement interface {
    Node
    statementNode()
}

type Instruction struct {
    Statement
    Mnemonic token.Token
    Operands []Operand
}

type Operand interface {
    Node
    operandNode()
}

type Register token.Token

type RegisterPair struct {
    Lower Register
    Upper Register
}

// Not a fan of this method of constraining types, but Go seems to 
// be doing it for their "go/ast" package, so there's probably no
// better way.
func (Instruction) statementNode() {}

func (Register)     operandNode() {}
func (RegisterPair) operandNode() {}

type SizeType byte

const (
    Byte SizeType = iota
    Word
)

