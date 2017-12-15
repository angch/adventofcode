package main

import (
	"fmt"
	"io/ioutil"
)

type Status int

const (
	InsideStream = Status(iota)
	EscapeNext
	InsideGarbage
	InsideGarbageEscapeNext
)

func main() {
	do("{}")
	do("{{{}}}")
	do("{{},{}}")
	do("{{{},{},{{}}}}")
	do("{<a>,<a>,<a>,<a>}")
	do("{{<ab>},{<ab>},{<ab>},{<ab>}}")
	do("{{<!!>},{<!!>},{<!!>},{<!!>}}")
	do("{{<a!>},{<a!>},{<a!>},{<ab>}}")
	f, _ := ioutil.ReadFile("input.txt")
	do(string(f))
}

func do(input string) {
	status := InsideStream
	level := 0
	score := 0
	garbageCount := 0

	for _, c := range input {
		switch status {
		case InsideStream:
			switch c {
			case '{':
				level++
			case '}':
				score += level
				level--
			case '!':
				status = EscapeNext
			case '<':
				status = InsideGarbage
			}
		case InsideGarbage:
			switch c {
			case '!':
				status = InsideGarbageEscapeNext
			case '>':
				status = InsideStream
			default:
				garbageCount++
			}
		case EscapeNext:
			status = InsideStream
		case InsideGarbageEscapeNext:
			status = InsideGarbage
		}
		//fmt.Println("status", c, status, level, score)
	}
	fmt.Println(input, "score", score, garbageCount)
}
