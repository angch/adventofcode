package main

import (
	"bytes"
	"os"
	"testing"
)

func BenchmarkOriginal(b *testing.B) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		b.Fatalf("could not read input file: %v", err)
	}

	r := bytes.NewReader(input)
	for n := 0; n < b.N; n++ {
		r.Reset(input)
		_, _ = day6io(r)
	}
}

func BenchmarkOptimized(b *testing.B) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		b.Fatalf("could not read input file: %v", err)
	}

	r := bytes.NewReader(input)
	for n := 0; n < b.N; n++ {
		r.Reset(input)
		_, _ = day6io2(r)
	}
}
