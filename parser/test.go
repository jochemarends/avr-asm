package main

import "fmt"

type A interface {
    Foo()
}

type B interface {
    C | D
}

type C struct {

}

func (C) Foo() {

}

type D struct {

}

func (D) Foo() {

}

func Bar(B) {

}

func main() {
    var c C
    Bar(c)
    fmt.Println("Ok")
}
