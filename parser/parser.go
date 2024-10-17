package parser

import (
    "fmt"
    "iter"
    "slices"
    "avrasm/ast"
    "avrasm/errors"
    "avrasm/token"
)

const (
    InvalidArgumentType errors.What = "invalid argument type"
)

type Parser struct {
    tokens []token.Token
    index  int
}

func New(tokens iter.Seq[token.Token]) *Parser {
    return &Parser{tokens: slices.Collect(tokens)}
}

func (p *Parser) Parse() (program []*ast.Instruction, errors []error) {
    for {
        instr, err := p.parseInstr()
        
        if err != nil {
            errors = append(errors, err)
        } else {
            program = append(program, instr)
        }
    }
}

func (parser *Parser) peekToken() (tok token.Token, err error) {
    if parser.index + 1 < len(parser.tokens) {
        tok = parser.tokens[parser.index]
    } else {
        err = fmt.Errorf("found no token to peek")
    }
    return
}

func (parser *Parser) currToken() (tok token.Token, err error) {
    if parser.index < len(parser.tokens) {
        tok = parser.tokens[parser.index]
    } else {
        err = fmt.Errorf("found no token")
    }
    return
}

func (parser *Parser) prevToken() (tok token.Token, err error) {
    if parser.index < len(parser.tokens) {
        tok = parser.tokens[parser.index]
    } else {
        err = fmt.Errorf("found no previous token")
    }
    return
}

func (parser *Parser) advanceToken() (tok token.Token, err error) {
    _, err = parser.peekToken()
    tok, _ = parser.currToken()
    parser.index++
    return
}

func (parser *Parser) matchToken(kinds ...token.Kind) bool {
    if parser.testToken(kinds...) {
        parser.advanceToken()
        return true
    }
    return false
}

func (parser *Parser) testToken(kinds ...token.Kind) bool {
    tok, err := parser.peekToken()

    if err != nil {
        return false
    }

    for _, kind := range kinds {
        if tok.Kind == kind {
            return true
        }
    }

    return false
}

func (parser *Parser) expectToken(expected token.Kind, message string) (tok token.Token, err error) {
    tok, err = parser.peekToken()

    if err == nil && tok.Kind != expected {
        err = fmt.Errorf("%s", message)
    }

    return
}

func (parser *Parser) parseInstr() (instr *ast.Instruction, err error) {
    mnemonic, err := parser.advanceToken()
    if err != nil {
        return
    }

    operands, err := parser.parseOperands()
    if err != nil {
        return
    }

    instr = &ast.Instruction{
        Mnemonic: mnemonic,
        Operands: operands,
    }

    return
}

func (parser *Parser) parseOperands() (ops []ast.Operand, err error) {
    for {
        var reg ast.Register
        reg, err = parser.parseRegister()

        if err != nil {
            return
        }

        ops = append(ops, reg)

        if parser.matchToken(token.Comma) {
            continue
        }

        break
    }

    return
}

func (parser *Parser) parseOperand() (op ast.Operand, err error) {
    tok, err := parser.peekToken()
    
    if err == nil {
        if tok.Kind == token.Register {
            return parser.parseRegister()
        }
        return parser.parseImmediate()
    }

    return
}

func (parser *Parser) parseRegister() (reg ast.Register, err error) {
    tok, err := parser.peekToken()

    if err != nil {
        err = fmt.Errorf("expected a token of kind 'register', but failed to read any token")
    } else if tok.Kind != token.Register {
        err = fmt.Errorf("expected a token of kind 'register', received '%v' instead", tok.Kind)
    } else {
        reg = ast.Register(tok)
    }

    return
}

func (parser *Parser) parseImmediate() (imm ast.Immediate, err error) {
    tok, err := parser.peekToken()

    if err == nil && tok.Kind == token.Number {
        imm = ast.Immediate(tok)
    }

    return
}

