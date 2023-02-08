package sample

import (
	"log"
	"strconv"
	"strings"

	"github.com/getgauge-contrib/gauge-go/gauge"
	m "github.com/getgauge-contrib/gauge-go/models"
)

var vowels = []string{"a", "e", "i", "o", "u"}

func numberOfVowels(word string) int {
	count := 0
	for _, char := range word {
		for _, vowel := range vowels {
			if string(char) == vowel {
				count++
			}
		}
	}
	return count
}

// Step implementations

var _ = gauge.Step("The word <word> has <number> vowels.", func(word string, number string) {
	expectedNumber := numberOfVowels(word)
	actualNumber, _ := strconv.Atoi(number)
	if expectedNumber != actualNumber {
		log.Printf("Expected number of vowels %d but got %d", expectedNumber, actualNumber)
	}
})

var _ = gauge.Step("Vowels in English language are <vowels>.", func(vowelsGiven string) {
	if vowelsGiven != strings.Join(vowels, "") {
		log.Printf("Expected vowels to be %s but got %s", vowelsGiven, strings.Join(vowels, ""))
	}
})

var _ = gauge.Step("Almost all words have vowels <table>", func(table *m.Table) {
	for _, row := range table.Rows {
		word := row.Cells[0]
		vowelCount := numberOfVowels(word)
		expectedVowelCount, _ := strconv.Atoi(row.Cells[1])
		if vowelCount != expectedVowelCount {
			log.Printf("Expected number of vowels for word '%s' to be %d but got %d", word, expectedVowelCount, vowelCount)
		}
	}
})
