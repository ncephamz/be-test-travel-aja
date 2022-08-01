package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// update input value here
	const input = 145050

	json, err := json.Marshal(rupiahDenomination(input))
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(json))
	}
}

func rupiahDenomination(input int) map[int]int {
	denomination := [10]int{
		100000,
		50000,
		20000,
		10000,
		5000,
		2000,
		1000,
		500,
		200,
		100,
	}

	result := make(map[int]int)

	for _, value := range denomination {
		if input == 0 {
			break
		}

		if input/value != 0 {
			result[value] += input / value
			input = input % value

			if input < 100 {
				input += 100
			}
		}
	}

	return result
}
