package main

import (
    "fmt"
    "strings"
)

var (
    ip = 0
    program = "10 20 + write"

    print = fmt.Printf
)

const (
    GOSTACK_PUSH = iota
    GOSTACK_WRITE = iota
    GOSTACK_PLUS = iota
    GOSTACK_OP_COUNT = iota
)

type gostack struct {
    name string
    curtok string
    nextok string
    code int64
}
var gs gostack

func printStruct() {
    print("Gostack: \n\tname: %s,\n\tcurtok: %s,\n\tcode: %d\n", gs.name, gs.curtok, gs.code)
}

func checkOp() {
    if (GOSTACK_OP_COUNT != 3) {
    print("Oops, you have unhandled operations.")
    }
}

// var status = "release"
var status = "devel"

func lexer(program string) {
    checkOp()
    var prog = strings.Split(program, " ")
    for (ip < len(prog)) {
        var code = prog[ip]
        print("%s\n", prog[ip])
        if (code == "write") {
            gs.name = "Write"
            gs.curtok = code
            gs.code = GOSTACK_WRITE  
            if (status == "devel") {
                printStruct()
            }
        } else if (code == "+") {
            gs.name = "Plus"
            gs.curtok = code
            gs.code = GOSTACK_PLUS
            if (status == "devel") {
                printStruct()
            }
        }
        ip += 1
    }
}


func main() {
    // print("Hello World!\n")
    lexer(program)
}
