package main

import (
    "avr-asm/lexer"
    "avr-asm/parser"
    "avr-asm/encoding"
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Two arguments are expected.
    // One for the program's name, and one for an input file.
    // Additional arguments are ignored.
    if len(os.Args) < 2 {
        fmt.Println("ERROR: no input file specified")
        os.Exit(1)
    }

    // Open the specified input file
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("ERROR:", err)
        os.Exit(1)
    }
    defer file.Close()
    
    var lines []string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    source := strings.Join(lines, "\n")

    for _, v := range lexer.Tokenize(source) {
        fmt.Println(v.Type)
    }

    l := lexer.New(source)
    fmt.Println("source:", source)
    parser := parser.New(l)
    program := parser.ParseProgram()

    for _, node := range program.Instructions {
        fmt.Println("nodne", node)
        fmt.Println(encoding.Encode(&node))
        //encoding.Encode(&node)
    }
}
