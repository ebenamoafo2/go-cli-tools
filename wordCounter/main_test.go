package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4

	res := count(b, false, false)
	if res != exp {
		t.Errorf("got %d, want %d", res, exp)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	exp := 3
	res := count(b, true, false)
	if res != exp {
		t.Errorf("got %d, want %d", res, exp)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 24
	res := count(b, false, true)
	if res != exp {
		t.Errorf("got %d, want %d", res, exp)
	}
}

func BenchmarkCountBytes(b *testing.B) {
	data := bytes.Repeat([]byte("word1 word2 word3\n"), 10000000000)
	for i := 0; i < b.N; i++ {
		count(bytes.NewReader(data), false, true)
	}
}
