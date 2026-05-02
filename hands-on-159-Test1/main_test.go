package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(2, 3)
	if result != 5 {
		t.Errorf("This is wrong, expected answer is %d but got %d", 5, result)
	}
}
