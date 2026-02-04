package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run word_counter.go input.txt")
	}
	filename := os.Args[1]

	if !strings.HasSuffix(filename, ".txt") {
		log.Fatal("Not a .txt file")
	}

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordMapping := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Fields(line)

		for _, word := range words {
			word = strings.ToLower(word)
			word = strings.Trim(word, ".,!?\\\"'():;")
			wordMapping[word] += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for word, count := range wordMapping {
		fmt.Printf("%v : %v \n", word, count)
	}

}
