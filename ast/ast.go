package ast

import (
    . "avr-asm/arch"
    "avr-asm/token"
    "strconv"
)

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

type Program struct {
    Instructions []Instruction
}

type Operand interface {
    Node
    Word() Word
    operandNode()
}

type Register token.Token

func (r Register) Word() Word {
    i, _ := strconv.Atoi(r.Text[1:])
    return Word(i)
}

type RegisterPair struct {
    Lower Register
    Upper Register
}

func (r RegisterPair) Word() Word {
    return r.Upper.Word()
}

// Not a fan of this method of constraining types, but Go seems to 
// be doing it for their "go/ast" package, so there's probably no
// better way.
func (Instruction) statementNode() {}

func (Register)     operandNode() {}
func (RegisterPair) operandNode() {}

