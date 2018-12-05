package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func BenchmarkFoo3(b *testing.B) {
	input, _ := ioutil.ReadFile("input.txt")
	input2 := strings.TrimSpace(string(input))
	for n := 0; n < b.N; n++ {
		foo3(input2)
	}
}

func BenchmarkFoo4(b *testing.B) {
	input, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		foo4(input)
	}
}
