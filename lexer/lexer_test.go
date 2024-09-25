package lexer

import (
    "strings"
    "testing"
    "avrasm/token"
)

func TestLexer(t *testing.T) {
    reader := strings.NewReader("str, r0, r1")

    lexer := New(reader)
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

    if len(tokens) != 6 {
        t.Fatalf("expected 6 tokens, received %v", len(tokens))
    }
}

