package parser

import (
	"slices"
	"strings"
	"testing"
	"avrasm/lexer"
	"avrasm/token"
)

func TestParser(t *testing.T) {
    reader := strings.NewReader("str r0, r1")

    lexer := lexer.New(reader)
    var tokens []token.Token

    for {
        tok, err := lexer.Scan()

        if err != nil {
            t.Fatal(err)
        }

        tokens = append(tokens, *tok)

        if tok.Kind == token.EOF {
            break
        }
    }

    parser := New(slices.Values(tokens))
    _, err := parser.parseInstr()

    if err != nil {
        t.Fatal(err)
    }
}

