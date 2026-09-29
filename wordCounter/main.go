package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	//Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	countBytes := flag.Bool("b", false, "Count bytes")

	// Parsing the flags provided by the user
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *countBytes))
}

func count(r io.Reader, countLines, countBytes bool) int {

	//A scanner is used to read text from a Reader(such as files)
	scanner := bufio.NewScanner(r)

	// If the count lines flag is not set, we want to count words so we define
	// the scanner split type to words (default is split by lines)
	switch {
	case countBytes:
		scanner.Split(bufio.ScanBytes)
	case countLines:
		scanner.Split(bufio.ScanLines)
	default:
		scanner.Split(bufio.ScanWords)
	}
	// Defining a counter
	wordCount := 0

	// For every word or line scanned, add 1 to the counter
	for scanner.Scan() {
		wordCount++
	}
	// Return the total
	return wordCount
}
