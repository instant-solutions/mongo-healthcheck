package main

import "testing"

func TestCheck(t *testing.T) {
	err := check("mongodb://localhost:27017")
	if err != nil {
		t.Fail()
	}
}
