package main

import "strings"

const dictNameMaxLength = 255

var dictContext *gfDict

type gfDict struct {
	name  string
	flags byte
	prev  *gfDict
	link  *gfDict
	code  uint32
}

var codeSection []int = []int{}

const (
	immediateFlag = 0x1
)

func initDictionary() {
	addMachinePrimitives()
}

func createDictionaryEntry(name string, codePointer uint32, code []int, flag byte) {
	entry := gfDict{name: name, flags: flag, code: codePointer}

	if dictContext == nil {
		dictContext = &entry
	} else {
		entry.prev, dictContext = dictContext, &entry
	}

	codeSection = append(codeSection, code...)
}

func addMachinePrimitives() {
	createDictionaryEntry("EXIT", I_EXIT, []int{I_EXIT}, 0)
	// arithmetic operations
	createDictionaryEntry("+", I_PLUS, []int{I_PLUS}, 0)
	createDictionaryEntry("-", I_MINUS, []int{I_MINUS}, 0)
	createDictionaryEntry("*", I_MULT, []int{I_MULT}, 0)
	createDictionaryEntry("/", I_DIV, []int{I_DIV}, 0)
	createDictionaryEntry("NEGATE", I_NEG, []int{I_NEG}, 0)
	createDictionaryEntry("MOD", I_MOD, []int{I_MOD}, 0)
	createDictionaryEntry("ABS", I_ABS, []int{I_ABS}, 0)
	createDictionaryEntry("MAX", I_MAX, []int{I_MAX}, 0)
	createDictionaryEntry("MIN", I_MIN, []int{I_MIN}, 0)
	// stack operations
	createDictionaryEntry("DUP", I_DUP, []int{I_DUP}, 0)
	createDictionaryEntry("DROP", I_DROP, []int{I_DROP}, 0)
	createDictionaryEntry(".", I_DOT, []int{I_DOT}, 0)
	createDictionaryEntry("SWAP", I_SWAP, []int{I_SWAP}, 0)
	createDictionaryEntry("OVER", I_OVER, []int{I_OVER}, 0)
	createDictionaryEntry("ROT", I_ROT, []int{I_ROT}, 0)
	createDictionaryEntry("NIP", I_NIP, []int{I_NIP}, 0)
	createDictionaryEntry("TUCK", I_TUCK, []int{I_TUCK}, 0)
	createDictionaryEntry("ROLL", I_ROLL, []int{I_ROLL}, 0)
	createDictionaryEntry("PICK", I_PICK, []int{I_PICK}, 0)
	// return stack operations
	createDictionaryEntry(">R", I_TO_R, []int{I_TO_R}, 0)
	createDictionaryEntry("R>", I_FROM_R, []int{I_FROM_R}, 0)
	createDictionaryEntry("R@", I_R_FETCH, []int{I_R_FETCH}, 0)
	// dict operations
	createDictionaryEntry("'", I_TICK, []int{I_TICK}, 0)
	createDictionaryEntry("EXECUTE", I_EXECUTE, []int{I_EXECUTE}, 0)

	// compiler operations
	createDictionaryEntry(":", I_COLON, []int{I_COLON}, 0)
	createDictionaryEntry("LITERAL", I_LITERAL, []int{I_LITERAL}, 0)
	createDictionaryEntry(";", I_SEMICOLON, []int{I_SEMICOLON}, immediateFlag)
	// helpers
	createDictionaryEntry(".S", I_DOTS, []int{I_DOTS}, 0)
	createDictionaryEntry("WORDS", I_WORDS, []int{I_WORDS}, 0)
	createDictionaryEntry("BYE", I_BYE, []int{I_BYE}, 0)
	createDictionaryEntry("SEE", I_SEE, []int{I_SEE}, 0)
	// system variables
	createDictionaryEntry("STATE", I_STATE, []int{I_STATE}, 0)
	createDictionaryEntry(">IN", I_GREATER_THAN_IN, []int{I_GREATER_THAN_IN}, 0)
	// variables
	createDictionaryEntry("@", I_DEREFERENCE, []int{I_DEREFERENCE}, 0)

	createDictionaryEntry("NOOP", I_NOOP, []int{I_NOOP}, 0)
}

func searchDictionary(name string) *gfDict {
	if dictContext == nil {
		return nil
	}

	nameCap := strings.ToUpper(name)
	currentEntry := dictContext

	for {
		if currentEntry.name == nameCap {
			return currentEntry
		}

		if currentEntry.prev == nil {
			return nil
		}

		currentEntry = currentEntry.prev
	}
}

func searchDictionaryByCode(code uint32) *gfDict {
	if dictContext == nil {
		return nil
	}

	currentEntry := dictContext

	for {
		if currentEntry.code == code {
			return currentEntry
		}

		if currentEntry.prev == nil {
			return nil
		}

		currentEntry = currentEntry.prev
	}
}

func isMachineWord(entry *gfDict) bool {
	return entry.code <= uint32(lastPrimitiveId)
}

func isUserDefinedWord(entry *gfDict) bool {
	return !isMachineWord(entry)
}
