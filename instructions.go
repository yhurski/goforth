package main

import (
	"fmt"
)

var lastPrimitiveId = I_NOOP

const (
	I_EXIT = iota
	I_PLUS
	I_MINUS
	I_MULT
	I_DIV
	I_NEG

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
	sum := operands[0] + operands[0]
	dataStack.push(sum)
}

func minusOp() {
	operands := dataStack.Popn(2)
	sum := operands[0] - operands[0]
	dataStack.push(sum)
}

func multOp() {
	operands := dataStack.Popn(2)
	sum := operands[0] * operands[0]
	dataStack.push(sum)
}

func divOp() {
	operands := dataStack.Popn(2)
	sum := operands[0] / operands[0]
	dataStack.push(sum)
}

func negateOp() {
	operand := dataStack.Pop()
	dataStack.push(-operand)
}

func tickOp() {
	// get next word from buffer
	// search for the word in the dict
	// push the word' xt onto stack
}

func executeOp() {
	execToken := dataStack.Pop()
	executePrimitive(execToken)
}

// func executeByToken(xt int) {

// }

func dotsOp() {
	fmt.Printf("S[dataStack.len()]:%v\n", dataStack)
}
