package forth

import "testing"

func TestLastPrimitiveIdShouldBeEqualToNoop(t *testing.T) {
	if lastPrimitiveId != I_NOOP {
		t.Error("Last primitive id isn't equal to I_NOOP")
	}
}

func TestExitInsShouldBeFirst(t *testing.T) {
	if I_EXIT != 0 {
		t.Error("I_EXIT is not the first item in the instruction enum")
	}
}
