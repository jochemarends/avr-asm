package parser

import (
    "avr-asm/ast"
    "avr-asm/lexer"
)

type Parser struct {
    l *lexer.Lexer
    currToken token.Token
    nextToken token.Token
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l}
}

func (p *Parser) parseStatement ast.Statement {
    tok := l.ReadToken()

    switch currToken.Type {
    case token.Mnemonic:
        return parseInstruction()
    }
}

func (p *Parser) parseInstruction ast.Instruction {
    
}


