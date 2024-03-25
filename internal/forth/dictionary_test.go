package forth

import (
	"testing"
)

func TestInitialDictContext(t *testing.T) {
	if dictContext != nil {
		t.Error("Initial dictContext is not null")
	}
}

func TestInitialCodeSection(t *testing.T) {
	if len(codeSection) != 0 {
		t.Errorf("Initial codeSection is not empty: %v", codeSection)
	}
}

func TestAddMachinePrimitivesAddsWords(t *testing.T) {
	totalMachineWords := 36
	addMachinePrimitives()

	if len(codeSection) != totalMachineWords {
		t.Errorf("There should be %d machine words, but %d now", totalMachineWords, len(codeSection))
	}

	t.Cleanup(resetDictionary)
}

func TestCreateDictionaryEntryAddsWord(t *testing.T) {
	currentNumberOfMachineWords := len(codeSection)
	createDictionaryEntry("+", I_PLUS, []int{I_PLUS}, 0)

	t.Run("Word bytecode was added to code section", func(t *testing.T) {
		if currentNumberOfMachineWords+1 != len(codeSection) {
			t.Errorf("Code section length before adding: %d, should be: %d", currentNumberOfMachineWords, len(codeSection))
		}
	})
	t.Run("Last word dict entry has `+` name", func(t *testing.T) {
		if dictContext.name != "+" {
			t.Errorf("Last dict word name: %s, should be: %s", dictContext.name, "+")
		}
	})
	t.Run("Last code in code section has I_PLUS bytecode", func(t *testing.T) {
		lastAddedByteCode := codeSection[len(codeSection)-1]
		if lastAddedByteCode != I_PLUS {
			t.Errorf("Last added bytecode: %d, should be: %d", lastAddedByteCode, I_PLUS)
		}
	})

	t.Cleanup(resetDictionary)
}

func resetDictionary() {
	dictContext = nil
	codeSection = nil
}
