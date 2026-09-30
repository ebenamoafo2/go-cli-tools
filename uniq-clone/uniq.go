package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
)

type Entry struct {
	Line  string
	Count int
}

func main() {
	showCount := flag.Bool("c", false, "show count before each line")
	onlyDups := flag.Bool("d", false, "show only lines that appear more than once")
	sortByCount := flag.Bool("s", false, "sort by count, highest first")
	flag.Parse()

	entries, err := countLines(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if *onlyDups {
		entries = filterDups(entries)
	}
	if *sortByCount {
		sortEntries(entries)
	}

	for _, e := range entries {
		if *showCount {
			fmt.Printf("%d %s\n", e.Count, e.Line)
		} else {
			fmt.Println(e.Line)
		}
	}
}

// countLines returns each unique line with its count, in first-seen order.
func countLines(r io.Reader) ([]Entry, error) {
	counts := make(map[string]int)
	var order []string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if counts[line] == 0 {
			order = append(order, line) // first time we see this line
		}
		counts[line]++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	entries := make([]Entry, 0, len(order))
	for _, line := range order {
		entries = append(entries, Entry{Line: line, Count: counts[line]})
	}
	return entries, nil
}

func filterDups(entries []Entry) []Entry {
	var out []Entry
	for _, e := range entries {
		if e.Count > 1 {
			out = append(out, e)
		}
	}
	return out
}

func sortEntries(entries []Entry) {
	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i].Count > entries[j].Count
	})
}
