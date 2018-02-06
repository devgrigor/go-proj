package sub

import "testing"

func TestDoub(t *testing.T) {
	var v float64
	v = Doub(15.3)

	if v != 30.6 {
		t.Error("Expected 30.6 but got ", v)
	}

	v = 4
	v = Doub(v)
	if  v != 8 {
		t.Error("Expected 8 but got ", v)
	}
}