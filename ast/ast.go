package ast

import "avrasm/token"

type Node interface {}

type Instruction struct {
    Mnemonic token.Token
    Operands []Operand
}

type Operand interface {}

type Register token.Token

type Immediate token.Token

