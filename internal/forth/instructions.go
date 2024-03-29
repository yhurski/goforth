package forth

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

var lastPrimitiveId = I_NOOP

var nestingLevel = 0

const (
	I_EXIT = iota
	// integer arithmetic operations
	I_PLUS
	I_MINUS
	I_MULT
	I_DIV
	I_NEG
	I_MOD
	I_ABS
	I_MAX
	I_MIN
	// stack operations
	I_DUP
	I_DROP
	I_DOT
	I_SWAP
	I_OVER
	I_ROT
	I_NIP
	I_TUCK
	I_ROLL
	I_PICK
	// return stack operations
	I_TO_R
	I_FROM_R
	I_R_FETCH
	// dict operations
	I_TICK
	I_EXECUTE

	// compiler operations
	I_COLON
	I_LITERAL
	I_SEMICOLON

	// helpers
	I_DOTS
	I_WORDS
	I_BYE
	I_SEE

	// system variables
	I_STATE
	I_GREATER_THAN_IN
	// variables

	I_DEREFERENCE

	I_NOOP // must always be last
)

func execute(execToken int) {
	fmt.Printf("execToken: %v\n", execToken)

	if execToken <= lastPrimitiveId {
		executePrimitive(execToken)
	} else {
		executeUserDefinedToken(execToken)
	}

	next()
}

func executeUserDefinedToken(execToken int) {
	// execToken is an offset of a word in the chain of instructions in
	// an user-defined word
	nestingLevel++
	ip = uint32(execToken)
	nextToken := codeSection[execToken]
	fmt.Printf("execToken: %v, nextToken: %v\n", execToken, nextToken)

	execute(nextToken)
}

func next() {
	fmt.Printf("next: nestingLevel: %v, ip: %v\n", nestingLevel, ip)
	if nestingLevel > 0 {
		ip++
		nextToken := codeSection[ip]
		execute(nextToken)
	}
}

func executePrimitive(execToken int) {
	// ip = uint32(execToken)

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
	case I_MOD:
		modOp()
	case I_ABS:
		absOp()
	case I_MAX:
		maxOp()
	case I_MIN:
		minOp()
	// stack operations
	case I_DUP:
		dupOp()
	case I_DROP, I_DOT:
		dropOp()
	case I_SWAP:
		swapOp()
	case I_OVER:
		overOp()
	case I_ROT:
		rotOp()
	case I_NIP:
		nipOp()
	case I_TUCK:
		tuckOp()
	case I_ROLL:
		rollOp()
	case I_PICK:
		pickOp()
	// return stack operations
	case I_TO_R:
		toROp()
	case I_FROM_R:
		fromROp()
	case I_R_FETCH:
		rFetch()
	// dict operations
	case I_TICK:
		tickOp()
	case I_EXECUTE:
		executeOp()
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
	case I_WORDS:
		wordsOp()
	case I_BYE:
		byeOp()
	case I_SEE:
		seeOp()

	// system variables
	case I_STATE:
		stateOp()
	case I_GREATER_THAN_IN:
		greaterThanIn()
	// variables
	case I_DEREFERENCE:
		dereferenceOp()

	case I_NOOP:
		break
	}
}

func exitOp() {
	if nestingLevel > 0 {
		nestingLevel--
	}

	returnAddress, err := returnStack.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}

	ip = uint32(returnAddress)
	nextToken := codeSection[ip]

	execute(nextToken)
}

func plusOp() {
	operands, err := dataStack.Popn(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	sum := operands[0] + operands[1]
	dataStack.Push(sum)
}

func minusOp() {
	operands, err := dataStack.Popn(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	sum := operands[0] - operands[1]
	dataStack.Push(sum)
}

func multOp() {
	operands, err := dataStack.Popn(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	sum := operands[0] * operands[1]
	dataStack.Push(sum)
}

func divOp() {
	operands, err := dataStack.Popn(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	if operands[1] == 0 {
		fmt.Println("Division by zero!")
		return
	}
	sum := operands[0] / operands[1]
	dataStack.Push(sum)
}

func negateOp() {
	operand, err := dataStack.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	dataStack.Push(-operand)
}

func modOp() {
	operands, err := dataStack.Popn(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	sum := operands[0] % operands[1]
	dataStack.Push(sum)
}

func absOp() {
	operand, err := dataStack.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	operand = int(math.Abs(float64(operand)))
	dataStack.Push(operand)
}

func maxOp() {
	operands, err := dataStack.Popn(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := max(operands[0], operands[1])
	dataStack.Push(result)
}

func minOp() {
	operands, err := dataStack.Popn(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := min(operands[0], operands[1])
	dataStack.Push(result)
}

func dupOp() {
	operand, err := dataStack.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	dataStack.Push(operand)
	dataStack.Push(operand)
}

func dropOp() {
	_, err := dataStack.Pop()

	if err != nil {
		fmt.Println(err)
		return
	}
}

func swapOp() {
	operands, err := dataStack.Popn(2)
	if err != nil {
		fmt.Println(err)
		return
	}

	operand1, operand2 := operands[0], operands[1]
	dataStack.Push(operand2)
	dataStack.Push(operand1)
}

func overOp() {
	operands, err := dataStack.Popn(2)
	if err != nil {
		fmt.Println(err)
		return
	}

	operand1, operand2 := operands[0], operands[1]
	dataStack.Push(operand1)
	dataStack.Push(operand2)
	dataStack.Push(operand1)
}

func rotOp() {
	operands, err := dataStack.Popn(3)
	if err != nil {
		fmt.Println(err)
		return
	}

	operand1, operand2, operand3 := operands[0], operands[1], operands[2]
	dataStack.Push(operand2)
	dataStack.Push(operand1)
	dataStack.Push(operand3)
}

func nipOp() {
	operands, err := dataStack.Popn(2)
	if err != nil {
		fmt.Println(err)
		return
	}

	operand1 := operands[0]
	dataStack.Push(operand1)
}

func tuckOp() {
	operands, err := dataStack.Popn(2)
	if err != nil {
		fmt.Println(err)
		return
	}

	operand1, operand2 := operands[0], operands[1]
	dataStack.Push(operand1)
	dataStack.Push(operand2)
	dataStack.Push(operand1)
}

func rollOp() {
	// 1 3 4 2 10 20 30 40 50
	// 4 roll .s
	// 1 3 4 2 20 30 40 50 10
	offset, err := dataStack.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}

	itemToRoll := 0

	if offset < dataStack.Len() {
		bottomOffset := dataStack.Len() - offset - 1
		itemToRoll = dataStack.Get(bottomOffset)
		copy((*dataStack)[bottomOffset:], (*dataStack)[bottomOffset+1:])
		*dataStack = (*dataStack)[:dataStack.Len()-1]
	}

	dataStack.Push(itemToRoll)
}

func pickOp() {
	// 1 3 4 2 10 20 30 40 50
	// 4 pick .s
	// 1 3 4 2 10 20 30 40 50 10
	offset, err := dataStack.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}

	itemToPick := 0

	if offset < dataStack.Len() {
		bottomOffset := dataStack.Len() - offset - 1
		itemToPick = dataStack.Get(bottomOffset)
	}

	dataStack.Push(itemToPick)
}

func toROp() {
	operand, err := dataStack.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}

	returnStack.Push(operand)
}

func fromROp() {
	var operand int = 0

	operand, _ = returnStack.Pop()
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
	word, ok := getWord()
	if !ok { // no name provided, ignore silently
		return
	}

	dictEntry := searchDictionary(word)
	if dictEntry == nil {
		return
	}

	dataStack.Push(int(dictEntry.code))
}

func executeOp() {
	execToken, err := dataStack.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}

	execute(int(execToken))
}

func colonOp() {
	if state == 1 {
		return // skip if already in compiling mode
	}

	word, _ := getWord()

	fmt.Printf("found word: %v\n", word)

	// add an entry to dict (not finished one though)
	createDictionaryEntry(strings.ToUpper(word), uint32(len(codeSection)), []int{}, smudgeFlag)

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
		fmt.Printf("literalOp ip: %v\n", ip)
		operand := codeSection[ip]
		dataStack.Push(operand)
		// ip++
	} else {

	}
}

func semicolonOp() {
	fmt.Printf("in semicolon\n")
	if state == 1 { // in compiling mode
		appendInsToCurrentDictEntry([]int{I_EXIT})
		dictContext.flags = dictContext.flags ^ smudgeFlag // remove the smudge flag
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

func wordsOp() {
	if dictContext != nil {

		for currentEntry, li := dictContext, 1; currentEntry != nil; currentEntry, li = currentEntry.prev, li+1 {
			fmt.Printf("%-10v", currentEntry.name)

			if li%5 == 0 {
				fmt.Print("\n")
			}
		}
	}
}

func byeOp() {
	dotsOp()
	os.Exit(0)
}

func seeOp() {
	word, ok := getWord()
	if !ok {
		fmt.Println("No name provided!")
		return
	}

	dictEntry := searchDictionary(word)
	if dictEntry == nil {
		fmt.Println("Name not found!")
		return
	}

	if isMachineWord(dictEntry) {
		fmt.Printf("(%v) - The word is defined as machine primitive.\n", word)
		return
	}

	definition := []string{":", word}

	for wordIp := dictEntry.code; codeSection[wordIp] != I_EXIT; wordIp++ {
		code := codeSection[wordIp]
		if code == I_LITERAL {
			wordIp++
			definition = append(definition, strconv.Itoa(codeSection[wordIp]))
		} else {
			dictEntry = searchDictionaryByCode(uint32(codeSection[wordIp]))
			if isUserDefinedWord(dictEntry) {
				// skip function call prolog
				definition = definition[:len(definition)-2]
			}
			definition = append(definition, strings.ToLower(dictEntry.name))
		}
	}

	definition = append(definition, ";")

	fmt.Println(strings.Join(definition, " "))
}

func stateOp() {
	address, _ := strconv.Atoi(fmt.Sprintf("%d", &state))
	dataStack.Push(address)
}

func greaterThanIn() {
	address, _ := strconv.Atoi(fmt.Sprintf("%d", &pIn))
	dataStack.Push(address)
}

func dereferenceOp() {
	address, err := dataStack.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}

	unsafePointer := unsafe.Pointer(uintptr(address))
	value := *(*int)(unsafePointer)

	dataStack.Push(value)
}

func appendInsToCurrentDictEntry(instructions []int) {
	if state == 1 {
		codeSection = append(codeSection, instructions...)
	} else {
		// produce an error
	}
}
