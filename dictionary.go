package main

const dictNameMaxLength = 255

var dictContext *gfDict

type gfDict struct {
	name  string
	flags byte
	prev  *gfDict
	link  *gfDict
	code  uint32
}

func initDictionary() {

}

func createDictionaryEntry(name string, code uint32, flag byte) {
	entry := gfDict{name: name, flags: flag, code: code}

	if dictContext == nil {
		dictContext = &entry
	} else {
		entry.prev, dictContext = dictContext, &entry
	}
}

func addMachinePrimitives() {
	createDictionaryEntry("EXIT", I_EXIT, 0)
	createDictionaryEntry("+", I_PLUS, 0)
	createDictionaryEntry("-", I_MINUS, 0)
	createDictionaryEntry("*", I_MULT, 0)
	createDictionaryEntry("/", I_DIV, 0)
	createDictionaryEntry("NEGATE", I_NEG, 0)
}
