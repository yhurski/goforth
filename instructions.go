package main

import (
	"fmt"
	"strconv"
)

var lastPrimitiveId = I_NOOP

const (
	I_EXIT = iota
	I_PLUS
	I_MINUS
	I_MULT
	I_DIV
	I_NEG
	I_DUP

	I_TICK
	I_EXECUTE

	I_DOTS

	I_NOOP // must always be last
)

func executePrimitive(execToken int) {
	switch execToken {
	case I_EXIT:
		exitOp()
	case I_PLUS:
		plusOp()
	case I_MINUS:
		minusOp()
	case I_MULT:
		multOp()
	case I_DIV:
		divOp()
	case I_NEG:
		negateOp()
	case I_DUP:
		dupOp()
	case I_TICK:
		executeOp()
	case I_EXECUTE:

	case I_DOTS:
		dotsOp()
	}
}

func exitOp() {

}

func plusOp() {
	operands := dataStack.Popn(2)
	sum := operands[0] + operands[1]
	dataStack.Push(sum)
}

func minusOp() {
	operands := dataStack.Popn(2)
	sum := operands[0] - operands[1]
	dataStack.Push(sum)
}

func multOp() {
	operands := dataStack.Popn(2)
	sum := operands[0] * operands[1]
	dataStack.Push(sum)
}

func divOp() {
	operands := dataStack.Popn(2)
	sum := operands[0] / operands[1]
	dataStack.Push(sum)
}

func negateOp() {
	operand := dataStack.Pop()
	dataStack.Push(-operand)
}

func dupOp() {
	operand := dataStack.Pop()
	dataStack.Push(operand)
	dataStack.Push(operand)
}

func tickOp() {
	// get next word from buffer
	// search for the word in the dict
	// push the word' xt onto stack
}

func executeOp() {
	execToken := dataStack.Pop()
	executePrimitive(int(execToken))
}

// func executeByToken(xt int) {

// }

func dotsOp() {
	result := ""
	for _, item := range *dataStack {
		result += strconv.Itoa(item) + " "
	}

	fmt.Printf("S[%d]:%v\n", dataStack.Len(), result)
}
