package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DoForth() {
	initDictionary()

	var inputStr string
	var err error
	var words []string
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>> ")

		inputStr, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		words = strings.Fields(inputStr)

		for _, word := range words {
			interpret(word)
		}
	}
}

func interpret(word string) {
	findOrCompile(word)
}

func findOrCompile(word string) {
	if !state {
		searchDictionary(word)
	} else {
		// in compile mode
	}

}
