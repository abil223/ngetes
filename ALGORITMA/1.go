package main

import (
	"fmt"
	"strings"
)

func reverseAlphaChars(inputString string) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabet = alphabet + strings.ToLower(alphabet)

	chars := []rune(inputString)
	alphaIndices := []int{}
	for i, char := range chars {
		if strings.ContainsRune(alphabet, char) {
			alphaIndices = append(alphaIndices, i)
		}
	}
	alphaChars := []rune{}
	for i := len(chars) - 1; i >= 0; i-- {
		if strings.ContainsRune(alphabet, chars[i]) {
			alphaChars = append(alphaChars, chars[i])
		}
	}

	for i, index := range alphaIndices {
		chars[index] = alphaChars[i]
	}

	returnString := string(chars)
	return returnString
}

func main() {
	inputString := "NEGIE1"
	fmt.Printf("input: %s\n", inputString)
}