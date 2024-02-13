package main

import (
	"fmt"
	"strconv"
	"strings"
)

var lastPrimitiveId = I_NOOP

const (
	I_EXIT = iota
	// integer arithmetic operations
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
	// return stack operations
	I_TO_R
	I_FROM_R
	I_R_FETCH

	I_TICK
	I_EXECUTE

	// compiler operations
	I_COLON
	I_LITERAL
	I_SEMICOLON

	// helpers
	I_DOTS

	I_NOOP // must always be last
)

func execute(execToken int) {
	if execToken <= lastPrimitiveId {
		executePrimitive(execToken)
	} else {
		executeUserDefinedToken(execToken)
	}
}

func executeUserDefinedToken(execToken int) {
	// ip = uint32(entry.code)

}

func executePrimitive(execToken int) {
	ip = uint32(execToken)

	switch execToken {
	case I_EXIT:
		exitOp()
	// integer arithmetic operations
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
	// return stack operations
	case I_TO_R:
		toROp()
	case I_FROM_R:
		fromROp()
	case I_R_FETCH:
		rFetch()
	// dict operations
	case I_TICK:
		executeOp()
	case I_EXECUTE:

	// compiler operations
	case I_COLON:
		colonOp()
	case I_LITERAL:
		literalOp()
	case I_SEMICOLON:
		semicolonOp()

	// helpers
	case I_DOTS:
		dotsOp()
	}
}

func exitOp() {
	returnAddress := returnStack.Pop()
	ip = uint32(returnAddress)
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

func toROp() {
	operand := dataStack.Pop()
	returnStack.Push(operand)
}

func fromROp() {
	var operand int = 0

	if returnStack.Len() > 0 {
		operand = returnStack.Pop()
	}

	dataStack.Push(operand)
}

func rFetch() {
	var operand int = 0

	if returnStack.Len() > 0 {
		operand = returnStack.Get(returnStack.Len() - 1)
	}

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

func colonOp() {
	fmt.Println("!!! colon")
	if state == 1 {
		return // skip if already in compiling mode
	}

	word, _ := getWord()
	// word, ok := getWord()
	// if !ok { // no name provided, ignore silently
	// 	return
	// }

	fmt.Printf("found word: %v\n", word)

	// add an entry to dict (not finished one though)
	createDictionaryEntry(strings.ToUpper(word), uint32(len(codeSection)), []int{}, 0)

	// switch machine mode to compiling
	state = 1
}

func literalOp() {
	if state == 0 {
		// dataStack.Pop()

		// word, ok := getWord()
		// if !ok { // no name provided, ignore silently
		// 	return
		// }

		ip++
		operand := codeSection[ip]
		dataStack.Push(operand)
	} else {

	}
}

func semicolonOp() {
	fmt.Printf("in semicolon\n")
	if state == 1 { // in compiling mode
		appendInsToCurrentDictEntry([]int{I_EXIT})
	}

	state = 0
}

func dotsOp() {
	result := ""
	for _, item := range *dataStack {
		result += strconv.Itoa(item) + " "
	}

	fmt.Printf("S[%d]:%v\n", dataStack.Len(), result)
}

func appendInsToCurrentDictEntry(instructions []int) {
	if state == 1 {
		codeSection = append(codeSection, instructions...)
	} else {
		// produce an error
	}
}
