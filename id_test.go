package gtools

import "testing"

func TestNewIDGen(t *testing.T) {
	idGen := NewIDGen()
	id := idGen.Next()
	need := uint64(1)
	if id != need {
		t.Errorf("TestNewIDGen error, wanted %d, got %d", need, id)
	}

	id = idGen.Next()
	id = idGen.Next()
	need = uint64(3)
	if id != need {
		t.Errorf("TestNewIDGen error, wanted %d, got %d", need, id)
	}
}
