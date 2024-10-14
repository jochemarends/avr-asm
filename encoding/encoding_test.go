package encoding

import (
	"fmt"
	"testing"
	"avrasm/ast"
	"avrasm/token"
)

func TestParser(t *testing.T) {
    instr := &ast.Instruction{
        Mnemonic: token.Token{
            Kind: token.Mnemonic,
            Lexeme: "mov",
        },
        Operands: []ast.Operand{
            ast.Register{
                Kind: token.Register,
                Lexeme: "r0",
            },
            ast.Register{
                Kind: token.Register,
                Lexeme: "r1",
            },
        },
    }

    encoded, err := EncodeInstr(instr)
    fmt.Println(encoded, err);
}
