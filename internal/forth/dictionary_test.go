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

func TestSearchByName(t *testing.T) {
	createDictionaryEntry("NEGATE", I_NEG, []int{I_NEG}, 0)
	createDictionaryEntry("MOD", I_MOD, []int{I_MOD}, 0)
	createDictionaryEntry("ABS", I_ABS, []int{I_ABS}, 0)

	t.Run("When name is in upper case", func(t *testing.T) {
		word := searchDictionary("MOD")
		if word == nil {
			t.Error("Word 'MOD' was added to dictionary but nil was found")
		}
	})

	t.Run("When name is in lower case", func(t *testing.T) {
		word := searchDictionary("abs")
		if word == nil {
			t.Error("Word 'ABS' was added to dictionary but nil was found")
		}
	})

	t.Cleanup(resetDictionary)
}

func TestSearchByCode(t *testing.T) {
	createDictionaryEntry("NEGATE", I_NEG, []int{I_NEG}, 0)
	createDictionaryEntry("MOD", I_MOD, []int{I_MOD}, 0)
	createDictionaryEntry("ABS", I_ABS, []int{I_ABS}, 0)

	t.Run("When code is 'I_NEG'", func(t *testing.T) {
		word := searchDictionaryByCode(I_NEG)
		if word == nil {
			t.Error("Code 'I_NEG' was added to dictionary but nil was found")
		}
	})

	t.Run("When name is 'I_ABS'", func(t *testing.T) {
		word := searchDictionaryByCode(I_ABS)
		if word == nil {
			t.Error("Code 'I_ABS' was added to dictionary but nil was found")
		}
	})

	t.Cleanup(resetDictionary)
}

func TestIsMachineWord(t *testing.T) {
	t.Run("'I_MULT' should be a machine word", func(t *testing.T) {
		createDictionaryEntry("*", I_MULT, []int{I_MULT}, 0)

		if !isMachineWord(dictContext) {
			t.Error("'I_MULT' should be a machine word")
		}
	})

	t.Run("'I_USER_DEFINED_WORD' should not be a machine word", func(t *testing.T) {
		I_USER_DEFINED_WORD := lastPrimitiveId + 1
		createDictionaryEntry("USER_DEFINED_WORD", uint32(I_USER_DEFINED_WORD), []int{I_USER_DEFINED_WORD}, 0)

		if isMachineWord(dictContext) {
			t.Error("'I_USER_DEFINED_WORD' should not be a machine word")
		}
	})

	t.Cleanup(resetDictionary)
}

func resetDictionary() {
	dictContext = nil
	codeSection = nil
}
