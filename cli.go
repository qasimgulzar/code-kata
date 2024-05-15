package main

import (
	"flag"
	"fmt"
	"log"
	"todo/cli/helpers/http"
)

const (
	IntTypeEven  = "Even"
	IntTypeOdd   = "Odd"
	BaseURL      = "https://jsonplaceholder.typicode.com"
	DefaultLimit = 20
)

func main() {
	cmd()
}

func determineStartAndStep(intType string) (start int, step int) {
	switch intType {
	case IntTypeEven:
		return 2, 2
	case IntTypeOdd:
		return 1, 2
	default:
		return 1, 1
	}
}

func processIntegers(start int, step int, limit int, client http.TodoClient) {
	for i := start; i <= limit; i += step {
		todo, err := client.Fetch(i)
		if err != nil {
			log.Println("Error fetching todo:", err)
			continue
		}

		status := "Incomplete"
		if todo.Completed == true {
			status = "Complete"
		}

		log.Println(fmt.Sprintf("%d %s has status %s", todo.Id, todo.Title, status))
	}
}
func cmd() {
	intType := flag.String("intType", IntTypeEven, "Type of integers to process (Even | Odd)")
	limit := flag.Int("limit", DefaultLimit, "Upper limit for the processed integers")
	maximum := 2 * *limit
	flag.Parse()
	start, step := determineStartAndStep(*intType)

	if *intType == IntTypeOdd {
		maximum = (2 * *limit) - 1
	}

	client := http.TodoClient{BaseURL: BaseURL}
	processIntegers(start, step, maximum, client)
}
