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
	createDictionaryEntry("+", I_PLUS, []int{I_PLUS}, 0)
	createDictionaryEntry("-", I_MINUS, []int{I_MINUS}, 0)
	createDictionaryEntry("*", I_MULT, []int{I_MULT}, 0)
	createDictionaryEntry("/", I_DIV, []int{I_DIV}, 0)
	createDictionaryEntry("NEGATE", I_NEG, []int{I_NEG}, 0)
	// '
	// execute
	createDictionaryEntry(".S", I_DOTS, []int{I_DOTS}, 0)
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
