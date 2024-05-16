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

func (p *Parser) readToken() {
    p.currToken = p.l.ReadToken()
    p.nextToken = p.l.ReadToken()
}

func (p *Parser) parseStatement() ast.Statement {
    switch p.currToken.Type {
    case token.Mnemonic:
        return p.parseInstruction()
    default:
        return nil
    }
}

func (p *Parser) expectNext(t token.Type) bool {
    if p.nextToken.Type == t {
        p.readToken()
        return true
    } else {
        p.nextTokenError(t)
        return false
    }
}

func (p *Parser) nextTokenError(t token.Type) {
    msg := fmt.Sprintf("ERROR: expected %s token, received %s token instead", t, p.nextToken.Type)
    p.errors = append(p.errors, msg)
}

func (p *Parser) parseInstruction() ast.Instruction {
    inst := ast.Instruction{Mnemonic: p.currToken}
    p.readToken()
    inst.Operands = p.parseOperands()
    return inst
}

func (p *Parser) parseOperands() []ast.Operand {
    operands := []ast.Operand{}

    if p.currToken.Type == token.Register {
        if p.nextToken.Type == token.Colon {
            operands = append(operands, p.parseRegisterPair())
        } else {
            operands = append(operands, ast.Register(p.currToken))
        }
    }

    return operands
}

func (p *Parser) parseRegisterPair() *ast.RegisterPair {
    pair := &ast.RegisterPair{Lower: ast.Register(p.currToken)}

    if !p.expectNext(token.Register) {
        return nil
    }

    pair.Upper = ast.Register(p.currToken)
    return pair
}

