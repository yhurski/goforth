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
	}

	for _, input := range test {
		t.Run(input.message, func(t *testing.T) {
			if input.variable != input.initValue {
				t.Errorf(input.errorMessage, input.variable)
			}
		})
	}
}

func TestInitialState(t *testing.T) {
	if state != 0 {
		t.Errorf("Initial state is not interpeting (0): %d", state)
	}
}

func TestInitialIp(t *testing.T) {
	if ip != 0 {
		t.Errorf("Initial ip is not 0: %d", ip)
	}
}

func TestInitialpIn(t *testing.T) {
	if pIn != 0 {
		t.Errorf("Initial pInt is not 0: %d", pIn)
	}
}

func TestInitialErrorCode(t *testing.T) {
	if pIn != 0 {
		t.Errorf("Initial errorCode is not 0: %d", errorCode)
	}
}
