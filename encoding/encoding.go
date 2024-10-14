package encoding

import (
	. "avrasm/arch"
	"avrasm/ast"
	"fmt"
	"slices"
	"strconv"
)

func EncodeInstr(instr *ast.Instruction) (encoded Word, err error) {
    mnemonic := instr.Mnemonic.Lexeme
    encoding, ok := encodings[mnemonic]

    if ok {
        encoded, err = encoding.Encode(instr)
    } else {
        err = fmt.Errorf("the %v instruction is not implemented", mnemonic)
    }

    return
}

var encodings = map[string]Instruction{
    "mov": {
        Encoding: "0010 11rd dddd rrrr",
        Operands: []Operand{
            {Symbol: 'd', Type: Register},
            {Symbol: 'r', Type: Register},
        },
    },
}

type Operand struct {
    Type       TypeEncoder
    Constraint ConstraintEncoder
    Symbol     rune
}

type Encoding string

func (enc *Encoding) Length() int {
    count := 0
    for _, r := range *enc {
        if r != ' ' {
            count++
        }
    }
    return count
}

func (enc *Encoding) Opcode() Word {
    var opcode Word
    length := enc.Length()

    idx := 0
    for _, r := range *enc {
        if r == '1' {
            pos := idx - (length - 1)
            opcode |= (1 << pos)
        } 

        if r != ' ' {
            idx++
        }
    }

    return opcode
}

func (enc *Encoding) CreateMask(symbol rune, word Word) Word {
    var mask Word
    var pos int

    for _, r := range slices.Backward([]rune(*enc)) {
        if r == symbol {
            bit := word & 1
            word >>= 1
            mask |= (bit << pos)
        } 

        if r != ' ' {
            pos++
        }
    }

    return mask
}

type Instruction struct {
    Encoding Encoding
    Operands []Operand
}

func (instr *Instruction) Encode(node *ast.Instruction) (encoded Word, err error) {
    if len(instr.Operands) != len(node.Operands) {
        panic("operand count does not match")
    }

    for i, op := range node.Operands {
        encoding := instr.Operands[i]

        var word Word
        word, err = encoding.Encode(op)

        if err != nil {
            return
        }

        encoded |= instr.Encoding.CreateMask(encoding.Symbol, word)
    }

    return
}

func (o *Operand) Encode(op ast.Operand) (word Word, err error) {
    word, err = o.Type(op)

    if err != nil {
        return
    }

    if o.Constraint != nil {
        word, err = o.Constraint(word)
    }

    return
}

type TypeEncoder func(o ast.Operand) (encoded Word, err error)

type ConstraintEncoder func(w Word) (encoded Word, err error)

func Register(op ast.Operand) (encoded Word, err error) {
    //r, _ := regexp.Compile("(?i)^r(3[0-1]|[01]?[0-9]|2[0-9])$")

    if reg, ok := op.(ast.Register); ok {
        var val int
        val, err = strconv.Atoi(reg.Lexeme[1:])
        encoded = Word(val)
        return
    }

    return 0, fmt.Errorf("expected a register")
}

func Min(value Word) ConstraintEncoder {
    return func(word Word) (encoded Word, err error) {
        if word >= value {
            encoded = word - value
        }
        return 0, fmt.Errorf("expected a value above or equal to %v, received %v", value, word)
    }
}

