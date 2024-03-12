package forth

import (
	"testing"
)

func TestInitialCoreValues(t *testing.T) {
	var test = []struct {
		message      string
		variable     any
		initValue    any
		errorMessage string
	}{
		{"state should be 0", state, 0, "Initial state is not interpeting (0): %d"},
		{"ip should be 0", ip, uint32(0), "Initial ip is not 0: %d"},
		{"pIn should be 0", pIn, 0, "Initial pIn is not 0: %d"},
		{"errorCode should be 0", errorCode, byte(0), "Initial errorCode is not 0: %d"},
	}

	for _, input := range test {
		t.Run(input.message, func(t *testing.T) {
			if input.variable != input.initValue {
				t.Errorf(input.errorMessage, input.variable)
			}
		})
	}
}
