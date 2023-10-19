package main

import "testing"

func Test_genSlice_checkSize(t *testing.T) {
	value := 100
	slice := genSlice(100)
	if value != len(slice) {
		t.Errorf("Expected slice lenght of %v, but go %v", value, len(slice))
	}
}

func Test_isOddOrEven_checkOddStatus(t *testing.T) {
	val := value(1)
	if val.isOddOrEven() != "Odd" {
		t.Errorf("Expected Odd")
	}
}

func Test_isOddOrEven_checkEvenStatus(t *testing.T) {
	val := value(3)
	if val.isOddOrEven() != "Even" {
		t.Errorf("Expected Even")
	}
}
