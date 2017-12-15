package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type layer struct {
	depth   int
	scanner int
	dir     int
	mod     int
}

func (l *layer) reset() {
	l.scanner = 0
	l.dir = 0
	if l.depth == 0 {
		l.scanner = -1
	}
	l.mod = (l.depth - 1) * 2
}

func (l *layer) tick() int {
	if l.depth == 0 {
		return -1
	}
	if l.dir == 0 {
		l.dir = 1
		l.scanner = 0
	}

	l.scanner += l.dir
	if l.scanner < 0 || l.scanner >= l.depth {
		l.dir = -l.dir
		l.scanner += l.dir
		l.scanner += l.dir
	}

	return l.scanner
}

func do1(inputs string) int {
	lines := strings.Split(inputs, "\n")
	layers := make([]layer, 0)
	for _, l := range lines {
		layer1 := layer{}
		i := 0
		fmt.Sscanf(l, "%d: %d", &i, &layer1.depth)
		for len(layers) < i {
			layers = append(layers, layer{})
		}
		layers = append(layers, layer1)
	}
	for i := range layers {
		layers[i].reset()
	}

	pos := 0
	alarm := 0
	picosec := 0
	for i := 0; pos < len(layers); i++ {
		fmt.Println("picosec bt", picosec, pos, layers)

		if pos >= 0 && pos < len(layers) && layers[pos].depth > 0 && layers[pos].scanner == 0 {
			fmt.Println("caught at", pos, "picosec", picosec)
			alarm += layers[pos].depth * pos
		}

		fmt.Println("picosec at", picosec, pos, layers)

		for j := range layers {
			layers[j].tick()
		}
		pos++
		picosec++
		fmt.Println("")
	}
	fmt.Println("alarm", alarm)
	return alarm
}

func do2b(inputs string) int {
	lines := strings.Split(inputs, "\n")
	layers := make([]layer, 0)
	for _, l := range lines {
		layer1 := layer{}
		i := 0
		fmt.Sscanf(l, "%d: %d", &i, &layer1.depth)
		for len(layers) < i {
			layers = append(layers, layer{})
		}
		layers = append(layers, layer1)
	}

	for l := range layers {
		layers[l].reset()
	}
a:
	for delay := 0; ; delay++ {
		picosec := delay
		for pos := 0; pos < len(layers); pos++ {
			if layers[pos].depth > 1 {
				if picosec%layers[pos].mod == 0 {
					continue a
				}
			}
			picosec++
		}
		fmt.Println("delay win", delay)
		return delay
	}
}

func main() {
	//input, err := ioutil.ReadFile("test.txt")
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	do2b(string(input))
	//log.Fatal(test1)
}
