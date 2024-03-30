package main

import (
	"fmt"
	"strings"
)

func longestWord(sentence string) (string, int) {
	words := strings.Fields(sentence) 
	longestWord := ""
	maxLength := 0

	for _, word := range words {
		currentLength := len(word)
		if currentLength > maxLength {
			longestWord = word
			maxLength = currentLength
		}
	}

	return longestWord, maxLength
}

func main() {
	sentence := "Saya sangat senang mengerjakan soal algoritma"
	longest, length := longestWord(sentence)
	fmt.Printf("Kata terpanjang: %s (%d characters)\n", longest, length)
}
