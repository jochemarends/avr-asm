package ast

import (
    . "avr-asm/arch"
    "avr-asm/token"
    "fmt"
    "strconv"
)

type Node interface {
    Token() token.Token
}

type Statement interface {
    Node
    statementNode()
}

type Instruction struct {
    Mnemonic token.Token
    Operands []Operand
}

func (instr *Instruction) Token() token.Token { return instr.Mnemonic }
func (instr *Instruction) statementNode() {}

type Program struct {
    Instructions []Instruction
}

type Operand interface {
    Node
    Word() Word
    operandNode()
}

type Immediate token.Token

func (imm *Immediate) Word() Word { 
    value, err := strconv.Atoi(imm.Text) 

    if err != nil {
        panic(fmt.Sprintf("parse error: %v", err))
    }

    return Word(value)
}

func (imm *Immediate) Token() token.Token { return token.Token(*imm) }
func (imm *Immediate) operandNode() {}

type Register token.Token

func (reg *Register) Word() Word {
    value, err := strconv.Atoi(reg.Text[1:])

    if err != nil {
        panic(fmt.Sprintf("parse error: %v", err))
    }

    return Word(value)
}

func (reg *Register) Token() token.Token { return token.Token(*reg) }
func (reg *Register) operandNode() {}

