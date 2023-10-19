package main

import "testing"

func Test_genSlice_checkSize(t *testing.T) {
	value := 100
	slice := genSlice(100)
	if value != len(slice) {
		t.Errorf("Expected slice lenght of %v, but go %v", value, len(slice))
	}
}
