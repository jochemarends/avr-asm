package token

type Kind string

const (
    Mnemonic Kind = "mnemonic"
    Register Kind = "register"
    Comma    Kind = ","
    EOF      Kind = "eof"
)

type Token struct {
    Kind   Kind
    Lexeme string
    Line   uint
}

