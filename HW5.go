package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Index struct {
	WordIndex map[string][]int
	Lines []string
}

type SearchResult struct{
	SearchWord string
	Results []ResultEntry
}

type ResultEntry struct {
	LineNum int
	Line string
}

func readLinesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func newIndex(lines []string) *Index {
	wordIndex := make(map[string][]int)
	
	for i, line := range lines {
		words := strings.Fields(line)
		for _, word := range words {
			wordIndex[word] = append(wordIndex[word], i)
		}
	}

	return &Index{WordIndex: wordIndex, Lines: lines}
}

func (index *Index) searchByWord(word string) SearchResult {
	var results []ResultEntry
	if indices, exists := index.WordIndex[word]; exists {
		for _, lineIndex := range indices {
			results = append(results, ResultEntry{LineNum: lineIndex, Line: index.Lines[lineIndex]})
		}
	}
	return SearchResult{SearchWord: word, Results: results}
}

func main() {
	filename := "text.txt" 
	lines, err := readLinesFromFile(filename)
	if err != nil {
		fmt.Printf("Помилка при зчитуванні файлу: %v\n", err)
		return
	}

	index := newIndex(lines)

	fmt.Print("Введіть слово для пошуку: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	searchWord := scanner.Text()

	results := index.searchByWord(searchWord)
	
	fmt.Printf("Результати пошуку за словом \"%s\":\n", searchWord)
	for _, result := range results.Results {
		fmt.Println(result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Помилка при зчитуванні введення: %v\n", err)
	}
}
