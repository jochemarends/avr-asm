package token

type Kind string

const (
    Mnemonic Kind = "mnemonic"
    Register Kind = "register"
    Number   Kind = "number"
    Comma    Kind = ","
    Newline  Kind = "\n"
    EOF      Kind = "eof"
)

type Token struct {
    Kind   Kind
    Lexeme string
    Line   uint
}

