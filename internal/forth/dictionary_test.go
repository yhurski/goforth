package forth

import "testing"

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
