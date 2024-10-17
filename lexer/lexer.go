package lexer

import (
    "bufio"
    "io"
    "regexp"
    "unicode"
    "avrasm/errors"
    "avrasm/token"
)

const (
    EndOfInput          errors.What = "end of input"
    UnexpectedCharacter errors.What = "unexpected character"
)

type Lexer struct {
    reader  *bufio.Reader
    buffer  []rune
    line    uint
    eof     bool
}

func isLetter(r rune) bool {
    return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isDigit(r rune) bool {
    return r >= '0' && r <= '9'
}

func isRegister(s string) bool {
    r := regexp.MustCompile("(?i)^r(3[0-1]|[01]?[0-9]|2[0-9])$")
	return r.MatchString(s)
}

func New(reader io.Reader) *Lexer {
    return &Lexer{reader: bufio.NewReader(reader)}
}

func (lexer *Lexer) peekRunes(n int) *rune {
    if n == 0 {
        panic("cannot peek zero character ahead")
    }

    for len(lexer.buffer) < n {
        r, _, err := lexer.reader.ReadRune()

        if err == nil {
            lexer.buffer = append(lexer.buffer, r)
        } else {
            return nil
        }
    }

    return &lexer.buffer[n - 1]
}

func (lexer *Lexer) peekRune() *rune {
    return lexer.peekRunes(1)
}

func (lexer *Lexer) readRune() *rune {
    r := lexer.peekRune()

    if r != nil {
        if *r == '\n' {
            lexer.line++
        }
        tail := lexer.buffer[1:]
        lexer.buffer = tail
    }

    return r
}

func (lexer *Lexer) putback(r rune) {
    // prepend
    lexer.buffer = append(lexer.buffer, r)

    if buflen := len(lexer.buffer); buflen > 0 {
        copy(lexer.buffer[1:], lexer.buffer)
        lexer.buffer[0] = r
    }
}

func (lexer *Lexer) readWhile(pred func(rune) bool) (text string) {
    for {
        r := lexer.peekRune()

        if r != nil {
            if pred(*r) {
                lexer.readRune()
                text += string(*r)
                continue
            }
        }
        break
    }

    return
}

func (lexer *Lexer) Scan() (*token.Token, *errors.Error) {
    r := lexer.readRune()

    if r == nil {
        // only emit one EOF token
        if !lexer.eof {
            lexer.eof = true
            return lexer.newToken(token.EOF, ""), nil
        }
        return nil, lexer.newError(EndOfInput)
    }

    switch *r {
    case ',':
        return lexer.newToken(token.Comma, string(*r)), nil
    default:
        if unicode.IsSpace(*r) {
            return lexer.Scan()
        }

        lexer.putback(*r)
        if isLetter(*r) {
            return lexer.scanName()
        } else {
            return lexer.scanNumber()
        }
    }
}

func (lexer *Lexer) scanName() (*token.Token, *errors.Error) {
    head := lexer.readRune()

    if head != nil && isLetter(*head) {
        tail := lexer.readWhile(func(r rune) bool {
            return isLetter(r) || isDigit(r)
        })

        name := string(*head) + tail
        kind := token.Mnemonic

        if isRegister(name) {
            kind = token.Register
        }

        return lexer.newToken(kind, name), nil
    }

    return nil, lexer.newError(UnexpectedCharacter)
}

func (lexer *Lexer) scanNumber() (*token.Token, *errors.Error) {
    text := lexer.readWhile(isDigit)
    
    if len(text) == 0 {
        return nil, lexer.newError(EndOfInput)
    } else {
        return lexer.newToken(token.Number, text), nil
    }
}

func (lexer *Lexer) newToken(kind token.Kind, lexeme string) *token.Token {
    return &token.Token{
        Kind:   kind,
        Lexeme: lexeme,
        Line:   lexer.line,
    }
}

func (lexer *Lexer) newError(kind errors.What) *errors.Error {
    return &errors.Error{
        What: kind,
        Line: lexer.line,
    }
}

