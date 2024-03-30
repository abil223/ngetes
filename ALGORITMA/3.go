package main

import (
	"fmt"
	"strings"
)

func detectStrings(inputArr []string, queryArr []string) ([]int, []string) {
	output := []int{}
	msg := []string{}

	for _, q := range queryArr {
		count := 0
		for _, string := range inputArr {
			if q == string {
				count++
			}
		}

		if count == 0 {
			msg = append(msg, fmt.Sprintf("kata %s tidak ada pada INPUT", q))
		} else {
			msg = append(msg, fmt.Sprintf("kata %s terdapat %d pada INPUT", q, count))
		}
		output = append(output, count)
	}

	return output, msg
}

func main() {
	input := []string{"xc", "dz", "bbb", "dz"}
	fmt.Printf("INPUT: %v\n", input)

	query := []string{"bbb", "ac", "dz"}
	fmt.Printf("QUERY: %v\n", query)

	output, message := detectStrings(input, query)
	fmt.Printf("result: %v\n", output)
	fmt.Printf("message: %v\n", strings.Join(message, ", "))
}