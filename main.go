package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func writeToFile(wordMapping map[string]int) error {
	file, err := os.OpenFile("mapping.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	i := 0

	for key, value := range wordMapping {
		fmt.Fprintf(writer, "key:%v value:%v\n", key, value)
		i += 1
		if i%1000 == 0 {
			if err := writer.Flush(); err != nil {
				return err
			}
		}
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	return nil

}

func processFile(fileName string, wordMapping map[string]int) error {
	fmt.Println(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

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
		return err
	}
	defer file.Close()

	return nil
}

func run(args []string) error {

	fileNames := make([]string, 0)

	wordMapping := make(map[string]int)

	defer writeToFile(wordMapping)

	for i := 1; i < len(args); i++ {
		if strings.HasSuffix(args[i], ".txt") {
			fileNames = append(fileNames, args[i])
		}
	}

	for index, filename := range fileNames {
		fmt.Println(index, filename)
	}

	for _, fileName := range fileNames {

		if err := processFile(fileName, wordMapping); err != nil {
			log.Printf("skipping %s: %v\n", fileName, err)
			continue
		}

	}

	return nil
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run word_counter.go input.txt")
	}

	if err := run(os.Args); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
