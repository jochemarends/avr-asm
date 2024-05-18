
type (
    // Constraints can change the way an operand gets 
    // encoded that's why it returns an integral
    Constraint func(ast.Operand) (uint16, error)

    Operand struct {
        Type       OperandType
        Constraint OperandConstraint
    }
)

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

