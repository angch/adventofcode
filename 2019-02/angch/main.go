package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var target = 19690720

func part1(prog []int, noun, verb int) int {
	program := make([]int, len(prog))
	copy(program, prog)
	ip := 0
	program[1] = noun
	program[2] = verb

	for {
		switch program[ip] {
		case 1:
			a, b, d := program[ip+1], program[ip+2], program[ip+3]
			program[d] = program[a] + program[b]
		case 2:
			a, b, d := program[ip+1], program[ip+2], program[ip+3]
			// log.Println(ip, a, b, d)
			program[d] = program[a] * program[b]
		case 99:
			return program[0]
		default:
			log.Fatal(program[ip])
		}
		ip += 4
	}

}

func parse(s string) []int {
	s1 := strings.Split(s, ",")
	o := make([]int, len(s1)*10)
	for k := range s1 {
		a, _ := strconv.Atoi(s1[k])
		o[k] = a
	}
	return o
}

func main() {
	if false {
		fmt.Println(part1(parse("1,0,0,0,99"), 12, 2))
		fmt.Println(part1(parse("2,3,0,3,99"), 12, 2))
		fmt.Println(part1(parse("2,4,4,5,99,0"), 12, 2))
		fmt.Println(part1(parse("1,1,1,4,99,5,6,0,99"), 12, 2))
	}
	if true {
		fileName := "input.txt"

		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Scan()
		prog := parse(scanner.Text())

		fmt.Println(part1(prog, 12, 2))

		// noun, verb := 0, 0
		found := make(chan int)
		for noun := 0; noun < 100; noun++ {
			for verb := 0; verb < 100; verb++ {
				go func(noun, verb int) {
					a := part1(prog, noun, verb)
					if a == 19690720 {
						found <- noun*100 + verb
					}
				}(noun, verb)
			}
		}
		log.Println(<-found)
	}

}
