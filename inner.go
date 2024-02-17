package main

import (
	"fmt"
	"strconv"

	"github.com/chzyer/readline"
)

func DoForth() {
	initDictionary()
	rlInstance, err := initReadline()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", codeSection)

	for {
		printPrompt(rlInstance)

		pIn = 0
		inputBuffer, err = rlInstance.Readline()
		if err != nil {
			fmt.Println(err)
			break
		}

		interpret()
	}
}

func interpret() {
	word, ok := getWord()

	for ; ok; word, ok = getWord() {
		findOrCompile(word)
	}
}

func findOrCompile(word string) {
	// fmt.Printf("word: %v\n", word)
	if state == 0 { // in interpretation mode
		dictEntry := searchDictionary(word)
		fmt.Printf("dictEntry: %v\n", dictEntry)
		if dictEntry != nil {
			// executePrimitive(int(dictEntry.code))
			execute(int(dictEntry.code))
		} else {
			number, err := strconv.Atoi(word)
			if err != nil {
				fmt.Println(err)
				return
			}

			dataStack.Push(number)
		}
	} else { // in compilation mode
		dictEntry := searchDictionary(word)
		// fmt.Printf("dictEntry: %v\n", dictEntry)
		if dictEntry != nil {
			fmt.Printf("FOUND: %v\n", dictEntry.name)
			if dictEntry.flags&immediateFlag == 1 {
				executePrimitive(int(dictEntry.code))
			} else {
				if dictEntry.code > uint32(lastPrimitiveId) { // user-defined word
					prologCode := []int{I_LITERAL, len(codeSection) + 4, I_TO_R, int(dictEntry.code)}
					appendInsToCurrentDictEntry(prologCode)
				} else {
					appendInsToCurrentDictEntry([]int{int(dictEntry.code)})
				}
			}
		} else {
			number, err := strconv.Atoi(word)
			if err != nil {
				fmt.Println(err)
				return
			}

			// compile a number with prepending LITERAL instruction
			appendInsToCurrentDictEntry([]int{I_LITERAL, number})
		}
	}

	fmt.Printf("%v\n", codeSection)

}

func initReadline() (*readline.Instance, error) {
	return readline.New(">>> ")
}

func printPrompt(rl *readline.Instance) {
	var stateCharacter rune

	if state == 1 { // compilation mode
		stateCharacter = 'c'
	} else { // interpretation mode
		stateCharacter = '>'
	}

	rl.SetPrompt(fmt.Sprintf("%c>> ", stateCharacter))
}
