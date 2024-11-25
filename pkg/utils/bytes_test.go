package utils

import "testing"

func TestBytesEquals(t *testing.T) {
	b1 := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b2 := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	if !BytesEquals(b1, b2) {
		t.Fatal("BytesEquals() failed!")
	}
}
