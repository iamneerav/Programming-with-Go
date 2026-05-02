package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(8, 2)
	if result != 10 {
		t.Errorf("This is wrong, expected answer is %d but got %d", 10, result)
	}
}

func TestSub(t *testing.T) {
	result := sub(8, 2)
	if result != 6 {
		t.Errorf("This is wrong, expected answer is %d but got %d", 6, result)
	}
}
