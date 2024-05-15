package main

import "testing"

func TestDetermineStartAndStepEven(t *testing.T) {
	start, step := determineStartAndStep("Even")
	if start != 2 || step != 2 {
		t.Fatalf("Wrong value for step or start.")
	}
}

func TestDetermineStartAndStepOdd(t *testing.T) {
	start, step := determineStartAndStep("Odd")
	if start != 1 || step != 2 {
		t.Fatalf("Wrong value for step or start.")
	}
}
