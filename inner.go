package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func DoForth() {
	initDictionary()

	fmt.Printf("%v\n", codeSection)

	// var inputStr string
	var err error
	// var words []string
	reader := bufio.NewReader(os.Stdin)

	for {
		printPrompt()

		// inputStr, err = reader.ReadString('\n')
		pIn = 0
		inputBuffer, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		// words = strings.Fields(inputStr)

		// for _, word := range words {
		// 	interpret(word)
		// }
		interpret()
	}
}

func interpret() {
	word, ok := getWord()

	for ; ok; word, ok = getWord() {
		findOrCompile(word)
	}
}

// func interpret(word string) {
// 	findOrCompile(word)
// }

func findOrCompile(word string) {
	// fmt.Printf("word: %v\n", word)
	if state == 0 { // in interpretation mode
		dictEntry := searchDictionary(word)
		// fmt.Printf("dictEntry: %v\n", dictEntry)
		if dictEntry != nil {
			executePrimitive(int(dictEntry.code))
		} else {
			number, err := strconv.Atoi(word)
			if err != nil {
				fmt.Println(err)
				return
			}

			dataStack.Push(number)
		}
	} else { // in compilation mode

	}

}

func printPrompt() {
	var stateCharacter rune

	if state == 1 { // compilation mode
		stateCharacter = 'c'
	} else { // interpretation mode
		stateCharacter = '>'
	}

	fmt.Printf("%c>> ", stateCharacter)
}
