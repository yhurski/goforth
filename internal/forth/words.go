package forth

import (
	"strings"
)

func getWord() (string, bool) {
	inputBufferLength := len(inputBuffer)
	var word strings.Builder
	if pIn >= inputBufferLength {
		return "", false
	}

	for ; pIn < inputBufferLength; pIn++ {
		if inputBuffer[pIn] == ' ' || inputBuffer[pIn] == '\n' {
			if word.Len() == 0 {
				// pIn++
				continue
			} else {
				break
			}
		}

		word.WriteByte(inputBuffer[pIn])
	}

	// fmt.Printf("word: %v\n", word.String())
	// fmt.Printf("pIn: %v\n", pIn)
	return word.String(), word.Len() > 0
}
