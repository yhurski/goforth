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
	// stack operations
	I_DUP
	I_DROP
	I_SWAP
	I_OVER

	I_TICK
	I_EXECUTE

	// compiler operations
	I_COLON

	// helpers
	I_DOTS

	I_NOOP // must always be last
)

func executePrimitive(execToken int) {
	switch execToken {
	case I_EXIT:
		exitOp()
	// integer arithmetic
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
	// stack operations
	case I_DUP:
		dupOp()
	case I_DROP:
		dropOp()
	case I_SWAP:
		swapOp()
	case I_OVER:
		overOp()
	// dict operations
	case I_TICK:
		executeOp()
	case I_EXECUTE:

	// compiler operations
	case I_COLON:
		colonOp()

	// helpers
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

func dropOp() {
	dataStack.Pop()
}

func swapOp() {
	firstOperand := dataStack.Pop()
	secondOperand := dataStack.Pop()
	dataStack.Push(firstOperand)
	dataStack.Push(secondOperand)
}

func overOp() byte {
	errCode := 0
	if dataStack.Len() < 2 {
		errCode = 1
		return byte(errCode)
	}
	operand := (*dataStack)[dataStack.Len()-2]
	dataStack.Push(operand)

	return byte(errCode)
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

func colonOp() {

}

func dotsOp() {
	result := ""
	for _, item := range *dataStack {
		result += strconv.Itoa(item) + " "
	}

	fmt.Printf("S[%d]:%v\n", dataStack.Len(), result)
}
