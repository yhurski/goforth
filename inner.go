package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DoForth() {
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
	// do something
}
