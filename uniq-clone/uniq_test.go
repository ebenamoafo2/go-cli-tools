package main

import (
	"bytes"
	"strings"
	"testing"
)

// equal compares two entry slices, treating nil and empty as the same.
func equal(a, b []Entry) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestCountLines(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entry
	}{
		{
			name:  "duplicates keep first-seen order",
			input: "apple\nbanana\napple\ncherry\nbanana\napple\n",
			want:  []Entry{{"apple", 3}, {"banana", 2}, {"cherry", 1}},
		},
		{
			name:  "empty input",
			input: "",
			want:  []Entry{},
		},
		{
			name:  "single line",
			input: "hello\n",
			want:  []Entry{{"hello", 1}},
		},
		{
			name:  "no trailing newline",
			input: "a\nb\na",
			want:  []Entry{{"a", 2}, {"b", 1}},
		},
		{
			name:  "blank lines are counted as lines",
			input: "a\n\na\n\n",
			want:  []Entry{{"a", 2}, {"", 2}},
		},
		{
			name:  "case sensitive",
			input: "Apple\napple\nAPPLE\n",
			want:  []Entry{{"Apple", 1}, {"apple", 1}, {"APPLE", 1}},
		},
		{
			name:  "all lines identical",
			input: "x\nx\nx\n",
			want:  []Entry{{"x", 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := countLines(bytes.NewBufferString(tt.input))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountLinesTooLong(t *testing.T) {
	// Scanner's default limit is 64KB per line.
	input := strings.Repeat("a", 70000) + "\n"
	_, err := countLines(bytes.NewBufferString(input))
	if err == nil {
		t.Error("expected an error for a line over 64KB, got nil")
	}
}

func TestFilterDups(t *testing.T) {
	tests := []struct {
		name  string
		input []Entry
		want  []Entry
	}{
		{
			name:  "keeps only counts above 1",
			input: []Entry{{"apple", 3}, {"banana", 2}, {"cherry", 1}},
			want:  []Entry{{"apple", 3}, {"banana", 2}},
		},
		{
			name:  "no duplicates returns nothing",
			input: []Entry{{"a", 1}, {"b", 1}},
			want:  []Entry{},
		},
		{
			name:  "empty input",
			input: []Entry{},
			want:  []Entry{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := filterDups(tt.input)
			if !equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortEntries(t *testing.T) {
	tests := []struct {
		name  string
		input []Entry
		want  []Entry
	}{
		{
			name:  "highest count first",
			input: []Entry{{"cherry", 1}, {"apple", 3}, {"banana", 2}},
			want:  []Entry{{"apple", 3}, {"banana", 2}, {"cherry", 1}},
		},
		{
			name:  "ties keep original order",
			input: []Entry{{"b", 2}, {"a", 2}, {"c", 5}},
			want:  []Entry{{"c", 5}, {"b", 2}, {"a", 2}},
		},
		{
			name:  "already sorted",
			input: []Entry{{"a", 3}, {"b", 1}},
			want:  []Entry{{"a", 3}, {"b", 1}},
		},
		{
			name:  "empty",
			input: []Entry{},
			want:  []Entry{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortEntries(tt.input)
			if !equal(tt.input, tt.want) {
				t.Errorf("got %v, want %v", tt.input, tt.want)
			}
		})
	}
}

func TestDupsThenSort(t *testing.T) {
	// Mirrors what main does with -d and -s together.
	entries, _ := countLines(bytes.NewBufferString("a\nb\nb\nc\nc\nc\n"))
	entries = filterDups(entries)
	sortEntries(entries)

	want := []Entry{{"c", 3}, {"b", 2}}
	if !equal(entries, want) {
		t.Errorf("got %v, want %v", entries, want)
	}
}

func BenchmarkCountLines(b *testing.B) {
	data := bytes.Repeat([]byte("apple\nbanana\ncherry\napple\n"), 50000)
	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		countLines(bytes.NewReader(data))
	}
}
