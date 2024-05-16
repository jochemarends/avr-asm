package parser

import (
    "avr-asm/ast"
    "avr-asm/lexer"
)

type Parser struct {
    l         *lexer.Lexer
    currToken token.Token
    nextToken token.Token
    errors:   []string
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l}
    p.nextToken()
    p.nextToken()
}

func (p *Parser) nextToken() token.Token {
    p.currToken := p.ReadToken
    p.nextToken := p.l.ReadToken()
}

func (p *Parser) parseStatement ast.Statement {
    switch p.currToken.Type {
    case token.Mnemonic:
        return p.parseInstruction()
    default:
        return nil
    }
}

func nextTokenIs(t token.Type) bool {
    return nextToken.Type == t
}

func (p *Parser) parseInstruction() ast.Instruction {
    return ast.Instruction{
        Mnemonic: p.currToken.Text()
        Operands: p.parseOperands()
    }
}

func (p *Parser) parseOperands []ast.Operand {
    operands := []ast.Operand{}

    switch p.nextToken.Type {
    case token.Register:
        operands = p.nextToken
    default:
        return operands
    }
}

func p *Pa

