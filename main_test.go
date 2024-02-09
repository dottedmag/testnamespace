package main

import "testing"

func TestTest(t *testing.T) {
	if 2+2 != 4 {
		t.Error("2+2 != 4")
	}
}
