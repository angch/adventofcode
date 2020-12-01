package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"testing"
)

var benchInput []int

func init() {
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	a := make([]int, 0)
	for scanner.Scan() {
		l := scanner.Text()
		i, _ := strconv.Atoi(l)
		a = append(a, i)
	}
	benchInput = a
}

func BenchmarkProcess2Clever(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process2clever(benchInput, 2020)
	}
}

func BenchmarkProcess2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process2(benchInput, 2020)
	}
}
