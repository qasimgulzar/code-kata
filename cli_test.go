package main

import (
	"testing"
	"todo/cli/helpers/http"
)

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

func TestFetchTodo(t *testing.T) {
	BaseURL := "https://jsonplaceholder.typicode.com"
	client := http.TodoClient{BaseURL: BaseURL}
	todo, err := client.Fetch(1)
	if err != nil {
		t.Fatalf("Error fetching todo.")
	}

	if todo.Id != 1 {
		t.Fatalf("Error fetching correct todo.")
	}
}
