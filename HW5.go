package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func indexText(lines []string) map[string][]int {
	wordIndex := make(map[string][]int)
	
	for i, line := range lines {
		words := strings.Fields(line)
		for _, word := range words {
			wordIndex[word] = append(wordIndex[word], i)
		}
	}

	return wordIndex
}

func searchByWord(word string, wordIndex map[string][]int, lines []string) []string {
	var results []string
	if indices, exists := wordIndex[word]; exists {
		for _, index := range indices {
			results = append(results, lines[index])
		}
	}
	return results
}

func main() {
	filename := "text.txt" 
	lines, err := readLinesFromFile(filename)
	if err != nil {
		fmt.Printf("Помилка при зчитуванні файлу: %v\n", err)
		return
	}

	wordIndex := indexText(lines)

	fmt.Print("Введіть слово для пошуку: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	searchWord := scanner.Text()

	results := searchByWord(searchWord, wordIndex, lines)
	
	fmt.Printf("Результати пошуку за словом \"%s\":\n", searchWord)
	for _, result := range results {
		fmt.Println(result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Помилка при зчитуванні введення: %v\n", err)
	}
}
