package main

var lastPrimitiveId = I_NOOP

const (
	I_EXIT = iota
	I_PLUS
	I_MINUS
	I_MULT
	I_DIV
	I_NEG

	I_NOOP // must always be last
)

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
