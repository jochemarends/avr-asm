package lexer

import (
    "avr-asm/token"
    "slices"
)

type Lexer struct {
    input        string
    readPosition int
}

func New(input string) *Lexer {
    return &Lexer{input: input}
}

func (l *Lexer) peekChar() byte {
    if l.readPosition < len(l.input) {
        return l.input[l.readPosition]
    }
    return 0
}

func (l *Lexer) readChar() byte {
    ch := l.peekChar()

    if ch != 0 {
        l.readPosition++
    }

    return ch
}

func (l *Lexer) unget() {
    l.readPosition = max(l.readPosition - 1, 0)
}

// Does not skip the '\n' character as it has its own token type.
func (l *Lexer) skipSpace() {
    spaces := []byte(" \t\v\f\r")

    for {
        if ch := l.peekChar(); slices.Contains(spaces, ch) {
            l.readChar()
        } else {
            break
        }
    }
}

func (l *Lexer) ReadToken() token.Token {
    var tok token.Token
    l.skipSpace()

    switch ch := l.readChar(); ch {
    case ',':
        tok = token.Token{Type: token.Comma, Text: string(ch)}
    case ':':
        tok = token.Token{Type: token.Colon, Text: string(ch)}
    case ' ', '\t', '\f', '\r':
        tok = token.Token{Type: token.Space, Text: string(ch)}
    case '\n':
        tok = token.Token{Type: token.EOL, Text: string(ch)}
    case 0:
        tok = token.Token{Type: token.EOF}
    default:
        if isLetter(ch) {
            l.unget()
            tok.Type = token.Symbol
            tok.Text = l.readSymbol()
            
            if token.IsMnemonic(tok.Text) {
                tok.Type = token.Mnemonic
            }

            if token.IsRegister(tok.Text) {
                tok.Type = token.Register
            }

        } else {
            tok.Type = token.Illegal
            tok.Text = string(ch)
        }
    }

    return tok
}

func (l *Lexer) readSymbol() string {
    startPosition := l.readPosition

    // First character can't be a number.
    if ch := l.peekChar(); isLetter(ch) {
        l.readChar()
        for {
            if ch := l.peekChar(); isLetter(ch) || isDigit(ch) {
                l.readChar()
            } else {
                break
            }
        }
    }

    return l.input[startPosition:l.readPosition]
}

func isLetter(ch byte) bool {
    isLower := ch >= 'a' && ch <= 'z'
    isUpper := ch >= 'A' && ch <= 'Z'
    return isLower || isUpper || ch == '_'
}

func isDigit(ch byte) bool {
    return ch >= '0' && ch <= '9'
}

func Tokenize(input string) []token.Token {
    var tokens []token.Token

    l := New(input)
    for {
        tok := l.ReadToken();
        tokens = append(tokens, tok)

        if tok.Type == token.EOF {
            break
        }
    }

    return tokens
}

