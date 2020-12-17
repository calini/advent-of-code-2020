package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const target = 2020

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <input.txt>\n", os.Args[0])
	}

	// Read file
	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Cannot open %s\n", os.Args[1])
	}
	defer inputFile.Close()

	// Create map
	numbersMap := make(map[int]struct{})

	// Read line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		// Convert line to int
		num, err := strconv.Atoi(scanner.Text())

		// Check for complement, if found
		if _, has := numbersMap[target - num]; has {
			// Output product
			log.Println(num * (target - num))

			// Found it, finish
			os.Exit(0)
		}

		// Add number to map
		if err != nil {
			log.Fatalf("Cannot convert line to number [%s]\n", scanner.Text())
		}
		numbersMap[num] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}



}
