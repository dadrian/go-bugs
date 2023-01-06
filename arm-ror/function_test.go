package arm_ror

import (
	"testing"
)

func TestShift(t *testing.T) {
	in := uint64(0x80000001)
	r := rotate16(in)
	expected := uint64(0x0001000000008000)
	if r != expected {
		t.Errorf("expected %x, got %x", expected, r)
	}
}
