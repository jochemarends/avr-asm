package object

type ObjectType

semantics.Analyse(program)

evaluator.Analyse(program)

encoder.Encode(program)

type Constraint func(ast.Node) bool

package semantics

type Instruction {
    Opcode  int
}

operands := map[opcode.OpcodeType][]Operand{
    ADD, {New(ByteRegister, Between(R10, R16)},
    ADD, {New(ByteRegister, Between(R10, R16),New(ByteRegister, Between(R10, R16)},
    ADD, {New(Between(R10, R16), New(Between(R10, R16))},
    ADD, {New(WordRegister, ))},
}


code.Operands[ADD]
code.Opcode[ADD]

 := "10101 10101 10101 rrrr"


type Operand struct {
    Type        OperandType
    Constraints []Constraint
}

func New(t OperandType, c ...Constraint) {

}

New(ByteRegister, Between(R10, R16))

type OperandType int


ADD = {NewOperand(ByteRegister, LowerHalf), NewOperand(ByteRegister, LowerHalf)}

const (
    Immediate OperandType = iota
    ByteRegister 
    WordRegister
)

