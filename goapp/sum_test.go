package main

import "testing"

func testSum(t *testing.T) {
	result := sum(2,3)
	if (result != 5) {
		t.Errorf("Inavlid! \n result: %d -> expected: %d", result, 5)
	}
}