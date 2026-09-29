package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

//grep clone
//./grep -i "error" < app.log prints matching lines.
//Uses: bufio.ScanLines, strings.Contains, flags -i (ignore case) and -n (line numbers).
//Tip: keep func search(r io.Reader, pattern string, ignoreCase bool) []string so it is easy to test with bytes.NewBufferString.

func main() {
	ignoreCase := flag.Bool("i", false, "ignore match")
	lineNumbers := flag.Bool("n", false, "line numbers only")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "usage: grep [-i] [-n] pattern")
		os.Exit(1)
	}

	pattern := flag.Arg(0)
	matches, err := search(os.Stdin, pattern, *ignoreCase)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, m := range matches {
		if *lineNumbers {
			fmt.Printf("%d:%s\n", m.Num, m.Text)
		} else {
			fmt.Println(m.Text)
		}
	}

}

type Match struct {
	Num  int
	Text string
}

func search(r io.Reader, pattern string, ignoreCase bool) ([]Match, error) {
	var matches []Match

	if ignoreCase {
		pattern = strings.ToLower(pattern)
	}
	scanner := bufio.NewScanner(r)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		haystack := line
		if ignoreCase {
			haystack = strings.ToLower(line)
		}
		if strings.Contains(haystack, pattern) {
			matches = append(matches, Match{Num: lineNum, Text: haystack})
		}
	}

	return matches, scanner.Err()
}
