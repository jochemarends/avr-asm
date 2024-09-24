package parser

import (
    "avr-asm/ast"
    "avr-asm/lexer"
    "avr-asm/token"
    "fmt"
)

type Parser struct {
    l         *lexer.Lexer
    currToken token.Token
    nextToken token.Token
    errors    []string
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l}
    p.readToken()
    p.readToken()
    return p
}

func (p *Parser) ParseProgram() *ast.Program {
    program := &ast.Program{}
    for p.currToken.Type != token.EOF {
        if instr := p.parseInstruction(); instr != nil {
            program.Instructions = append(program.Instructions, *instr)
        }
    }
    return program
}

func (p *Parser) readToken() {
    p.currToken = p.nextToken
    p.nextToken = p.l.ReadToken()
}

func (p *Parser) nextTokenError(t token.Type) {
    msg := fmt.Sprintf("ERROR: expected %s token, received %s token instead", t, p.nextToken.Type)
    p.errors = append(p.errors, msg)
}

func (p *Parser) parseStatement() ast.Statement {
    switch p.currToken.Type {
    case token.Mnemonic:
        return p.parseInstruction()
    default:
        return nil
    }
}

func (p *Parser) parseInstruction() *ast.Instruction {
    instr := &ast.Instruction{Mnemonic: p.currToken}
    p.readToken()
    instr.Operands = p.parseOperands()
    return instr
}

{ 
    Mnemonic: "ADD",
    Encoding: "0000 0001 dddd rrrr",
    Operands: [Register.Between(16, 31), Register.Between(16, 31)]
}

{ 
    Mnemonic: "ADD",
    Encoding: "0000 0001 dddd rrrr",
    Operands: [Register('d'), Register('r')]
}

{ 
    Mnemonic: "ORI",
    Encoding: "0110 KKKK dddd KKKK",
    Operands: [Register('d').Between(16, 23), Immediate('K').Bits(8)]
}

type OperandContraint func(int) int

Parse("ORI R16, 5")

func (p *Parser) parseOperands() []*ast.Operand {
    operands := []*ast.Operand{}

    p.readToken()

    for {
        if p.currToken.Type == token.EOL {
            return operands
        }

        if op := p.parseOperand(); op != nil {
            operands = append(operands, &op)
        } else {
            return nil
        }

        if token.

        break
    }

    Between(0, 16)


    for p.nextToken.Type == token.Comma {
        
    }

    if p.currToken.Type != token.Register {
        return nil
    }
    
    if p.currToken.Type == token.Register {
        if p.nextToken.Type == token.Colon {
            operands = append(operands, p.parseRegisterPair())
        } else {
            operands = append(operands, ast.Register(p.currToken))
        }
    }

    return operands
}

func (p *Parser) parseOperand() ast.Operand {
    if p.currToken.Type != token.Register {
        return nil
    }
    reg := ast.Register(p.currToken)
    return &reg
}

