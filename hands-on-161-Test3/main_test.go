package main

import "testing"

func TestLocation(t *testing.T) {

	current_location := location("Maldives")
	imaginary_location := ""

	if current_location != "I am in Maldives" {
		t.Errorf("This is wrong, expected answer is %s but got %s", current_location, imaginary_location)
	}

}
