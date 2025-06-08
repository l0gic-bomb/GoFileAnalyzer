package main

import (
	"testing"
)

func TestCleanWord(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{",,,", ""},
		{"hello, world", "helloworld"},
		{"hello,..", "hello"},
		{"hi,hi,hello!43242,woooor.ld", "hihihello43242woooorld"},
	}

	for _, test := range tests {
		result := cleanWord(&test.input)
		if result != test.expected {
			t.Errorf("cleanWord(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}

func TestUniqueWords(t *testing.T) {
	analyzer := &FileAnalyzer{
		frequencyWords: map[string]int{
			"hello": 1,
			"world": 2,
			"gooo":  1,
			"testt": 3,
			"test":  1,
			"cpp":   1,
		},
	}

	expected := 4
	analyzer.countUniqueWords()
	result := analyzer.uniqueWords
	if result != expected {
		t.Errorf("countUniqueWords() = %d, want %d", result, expected)
	}

	analyzer = &FileAnalyzer{}

	expected = 0
	analyzer.countUniqueWords()
	result = analyzer.uniqueWords
	if result != expected {
		t.Errorf("countUniqueWords() = %d, want %d", result, expected)
	}
}

func TestCountBytes(t *testing.T) {
	analyzer := &FileAnalyzer{}

	wordByBytes := make(map[string]int)
	wordByBytes["Hello"] = len("Hello")
	wordByBytes["Golang is forever"] = len("Golang is forever")
	wordByBytes[" "] = len(" ")
	wordByBytes["!!!!!tbertbre"] = len("!!!!!tbertbre")

	for word, bytes := range wordByBytes {
		analyzer.countBytes(word)
		if analyzer.bytes != bytes {
			t.Errorf("countBytes() = %d, want %d", analyzer.bytes, bytes)
		}
	}
}
