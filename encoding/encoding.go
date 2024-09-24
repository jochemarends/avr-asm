package encoding

import (
    . "avr-asm/arch"
    "avr-asm/ast"
    _ "avr-asm/parser"
    "fmt"
)

type (
    // Maps a numeric value to a new format, (e.g.) a value that can only reside in four states
    // can be represented using two bits.
    //
    // There's probably a better name for this, but I couldn't come up with one.
    Mapping func(Word) (Word, error)

    Opcode string
)

// Checks if a value is between two values (inclusive).
func Between(a, b Word) Mapping {
    minVal := min(a, b)
    maxVal := max(a, b)
    return func(w Word) (Word, error) {
        if w >= minVal && w <= maxVal {
            return w - minVal, nil
        }
        err := fmt.Errorf("expected a value from %v through %v, received %v instead", minVal, maxVal, w)
        return w, err
    }
}

// Returns its argument unchanged.
func Identity() Mapping {
    return func(w Word) (Word, error) {
        return w, nil
    }
}

// AVR Mnemonics
var Mnemonics = []string{
    "MOV",
}

// Constraints/mappings for each mnemonic.
//
// "MOV" accepts two arguments of any type. (It doesn't really, this is just for testing).
var Mappings = map[string][]Mapping{
    "MOV": {Identity(), Identity()},
}

// Opcodes for each mnemonic.
//
// Represented as string as can be found in the AVR instruction set manual.
var Opcodes = map[string]string{
    "MOV": "001011rdddddrrrr",
}

func Encode(node *ast.Instr) (Word, error) {
    // Error phrone, probably panics if entry exists.
    mappings := Mappings[node.Mnemonic.Text]
    opcode := Opcodes[node.Mnemonic.Text]

    // Check if instruction constains the correct amount of operands.
    if len(node.Operands) != len(mappings) {
        return 0, fmt.Errorf("expected %v operands, received %v instead", len(mappings), len(node.Operands))
    }

    var bin Word

    // Check if each operand can be mapped.
    for i, operand := range node.Operands {
        word, err := mappings[i](operand.Word())
        if err != nil {
            return word, err
        }
        bin |= Mask(opcode, 'r', word)
    }
    
    return bin, nil
}

func Mask(opcode string, symbol rune, word Word) Word {
    var mask Word

    for i, r := range opcode {
        if r == symbol {
            fmt.Println(i)
            bit := word & 1
            mask |= bit << (15 - i)
            word = word >> 1
        }
    }

    return mask
}

