package task

import "testing"

func TestSomthing(t *testing.T) {
	result := 1 + 1
	if result != 2 {
		t.Errorf("Expedted 2 bug got %d", result)
	}
}
