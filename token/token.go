package token

import "fmt"

type Token struct {
    Type Type
    Text string
}

type Type string

const (
    Mnemonic = "MNEMONIC"
    Register = "REGISTER"
    Symbol   = "SYMBOL"
    Comma    = ","
    Colon    = ":"
    Illegal  = "ILLEGAL"
    Space    = "SPACE"
    EOL      = "EOL"
    EOF      = "EOF"
)

func IsMnemonic(s string) bool {
    return s == "STR"
}

func IsRegister(s string) bool {
    var r rune
    var n int

    if _, err := fmt.Sscanf(s, "%c%d", &r, &n); err != nil {
        return false
    }

    if r != 'r' && r != 'R' {
        return false
    }

    if !(n >= 0 && n <= 31) {
        return false
    }

    return true
}

