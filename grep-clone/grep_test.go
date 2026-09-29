package main

import (
	"bytes"
	"testing"
)

func TestSearch(t *testing.T) {
	input := "Error: disk full\nall good\nerror again\n"

	got, _ := search(bytes.NewBufferString(input), "error", false)
	if len(got) != 1 {
		t.Errorf("case sensitive: got %d, want 1", len(got))
	}

	got, _ = search(bytes.NewBufferString(input), "error", true)
	if len(got) != 2 {
		t.Errorf("ignore case: got %d, want 2", len(got))
	}
	if got[1].Num != 3 {
		t.Errorf("line number: got %d, want 3", got[1].Num)
	}
}
