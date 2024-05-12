package instructions

type Instruction struct {
    Mnemonic string
    Opcode   int16
    Operands 
    Encoding
}

encoding := "000111rdddddrrrr"
operands := {Pair[Register]}

func parseRegisterPair()

{RegisterPair, RegisterPair}

const (
    Register := iota
    RegisterPair
    Immediate
    RegsterPair
)

type Opcode struct {
    int16 Value
    int16 Mask
}

opcode.MOVW

opcode.ST

operands.Find(opcode.MOVW)

opcode := "0000 0001 dddd rrrr"

const (
    MOVW := "0000 0001 dddd rrrr"
)

const (
    MOVW := "0000 0001 dddd rrrr"
)

const (
    MOVW := "0000 0001 dddd rrrr
)

bin := encode(opcode.MOVW, 0, 1)
