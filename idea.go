
type (
    // Constraints can change the way an operand gets 
    // encoded that's why it returns an integral
    Constraint func(ast.Operand) (uint16, error)
    
    Constraint func(uint16) (uint16, error)

    Operand struct {
        Type       OperandType
        constraints []Constraint
    }
)

func (o *Operand) Encode(node *ast.Operand) (val uint16, err error) {
    val := node.Int()
    for _, mapping := range mappings {
        val, err = mapping(value)
        if error == nil {
            return
        }
    }
}

func Encode(node *ast.Instruction) (val uint16, err error) {
    params := parameters[ast.Mnemonic]

    if len(params) != len node.Operands {
        // ERROR
    }

    var operands []uint16

    for _, operand := range node.Operands {
        // CHECK EACH
    }

    for _, operand := range node.Operands {
        ch := operand.Char()
        for idx, bit := range opcode {
            if bit == ch {
                setBit(val, idx)
            }
        }
    }
}

func Between(a, b uint16) Constraint {
    min := min(a, b)
    max := max(a, b)

    return func(o ast.Operand) (uint16, error) {
        if o.Int() >= min && o.Int() <= max {
            return o.Int() - min
        }
        return someError
    }
}

opcode := "1010 1010 rrrr dddd"
operands := {NewOperand('r', Between(R10, R16))}



opcode.Parse("1010 1010 rrr ddd", map[rune]uint16)

func Parse(opcode string, operands map[rune]uint16) uint16 {
    for _, o := range operands {
         
    }
}

const (
    ByteRegister OperandType = iota
    WordRegister
)
