package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

type FileAnalyzer struct {
	lineCount      int
	wordCount      int
	bytes          int
	frequencyWords map[string]int
}

func cleanWord(word *string) string {
	if word == nil {
		return ""
	}
	return strings.Map(func(symbol rune) rune {
		if unicode.IsLetter(symbol) || unicode.IsDigit(symbol) {
			return symbol
		}
		return -1
	}, *word)
}

func (analyzer *FileAnalyzer) printTopTenWords() {
	type wordCount struct {
		word  string
		count int
	}
	var wordsCounts []wordCount
	for key, value := range analyzer.frequencyWords {
		wordsCounts = append(wordsCounts, wordCount{key, value})
	}

	sort.Slice(wordsCounts, func(left, right int) bool {
		return wordsCounts[left].count > wordsCounts[right].count
	})

	for i := 0; i < 10 && i < len(wordsCounts); i++ {
		fmt.Printf("count %d for word %s \n", wordsCounts[i].count, wordsCounts[i].word)
	}
}

func (analyzer *FileAnalyzer) countFrequencyWords(allWords []string) {
	if analyzer.frequencyWords == nil {
		analyzer.frequencyWords = make(map[string]int)
	}
	for _, word := range allWords {
		cleanedWord := cleanWord(&word)
		if cleanedWord != "" {
			analyzer.frequencyWords[cleanedWord] += 1
		}
	}

	analyzer.printTopTenWords()
}

func (analyzer *FileAnalyzer) countBytes(text string) {
	var bytes int
	bytes = len(text)
	analyzer.bytes = bytes
}

func (analyzer *FileAnalyzer) analyzeText(text string) {
	analyzer.lineCount = strings.Count(text, "\n") + 1
	allWords := strings.Fields(text)
	analyzer.wordCount = len(allWords)
	analyzer.countFrequencyWords(allWords)
	analyzer.countBytes(text)
	fmt.Printf("Text has %d lines and %d words and %d bytes\n", analyzer.lineCount, analyzer.wordCount, analyzer.bytes)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("You don't get programm file's name\n")
		return
	}

	filename := os.Args[1]
	text, error := os.ReadFile(filename)
	if error != nil {
		fmt.Printf("Error reading file : %v\n", error)
		return
	}

	analyzer := &FileAnalyzer{
		frequencyWords: make(map[string]int),
	}
	analyzer.analyzeText(string(text))
}
