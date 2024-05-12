package ast

import "avr-asm/token"

type Node interface {
    Token() token.Token
}

type Statement interface {
    Node
}

type Instruction interface {
    Node
}

type Register interface {
    Node
    Size() SizeType
}

type Operand interface {
    Node
}

type Register interface {
    
}

type RegisterPair interface {
    
}

type Number interface {

}



type X {
    
}



type SizeType byte

const (
    Byte SizeType iota 
    Word
)

