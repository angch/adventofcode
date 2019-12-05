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

func part1(prog []int, input int) int {
	program := make([]int, len(prog))
	copy(program, prog)
	ip := 0
	output := make([]int, 0)

	for {
		opcode := program[ip] % 100
		code := program[ip] / 100

		param1 := code % 10
		param2 := (code / 10) % 10
		param3 := (code / 100) % 10
		if param3 > 0 {
			log.Println("param3 in usez!")
		}
		// log.Println(ip, opcode)

		switch opcode {
		case 1:
			a, b, d := program[ip+1], program[ip+2], program[ip+3]
			// log.Println(ip, "ADD", a, b, d, param1, param2, param3)
			if param1 == 0 {
				a = program[a]
			}
			if param2 == 0 {
				b = program[b]
			}

			program[d] = a + b
			ip += 4
		case 2:
			a, b, d := program[ip+1], program[ip+2], program[ip+3]
			// log.Println(ip, "MULT", a, b, d, param1, param2, param3)
			if param1 == 0 {
				a = program[a]
			}
			if param2 == 0 {
				b = program[b]
			}
			program[d] = a * b
			ip += 4
		case 3:
			d := program[ip+1]
			// log.Println(ip, "IN", d, param1)
			program[d] = input
			ip += 2
		case 4:
			d := program[ip+1]

			if param1 == 0 {
				d = program[d]
			}

			// log.Println(ip, "OUT", d, param1)
			output = append(output, d)
			ip += 2
		case 5:
			a, d := program[ip+1], program[ip+2]
			// log.Println(ip, "JNZ", a, d, param1, param2)
			if param1 == 0 {
				a = program[a]
			}
			if param2 == 0 {
				d = program[d]
			}
			if a != 0 {
				// log.Println("t", a)
				ip = d
			} else {
				// log.Println("f")
				ip += 3
			}
		case 6:
			a, d := program[ip+1], program[ip+2]
			// log.Println(ip, "JZ", a, d, param1, param2)
			if param1 == 0 {
				a = program[a]
			}
			if param2 == 0 {
				d = program[d]
			}
			if a == 0 {
				// log.Println("t")
				ip = d
			} else {
				// log.Println("f")
				ip += 3
			}
		case 7:
			a, b, d := program[ip+1], program[ip+2], program[ip+3]
			// log.Println(ip, "CMPL", a, b, d, param1, param2)
			if param1 == 0 {
				a = program[a]
			}
			if param2 == 0 {
				b = program[b]
			}
			if a < b {
				program[d] = 1
			} else {
				program[d] = 0
			}
			ip += 4
		case 8:
			a, b, d := program[ip+1], program[ip+2], program[ip+3]
			// log.Println(ip, "CMP", a, b, d, param1, param2)
			if param1 == 0 {
				a = program[a]
			}
			if param2 == 0 {
				b = program[b]
			}
			if a == b {
				// log.Println("t")
				program[d] = 1
			} else {
				// log.Println("f")
				program[d] = 0
			}
			ip += 4
		case 99:
			log.Println("x", output)
			return program[0]
		default:
			log.Fatal("fatal", program[ip])
		}

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
	if true {
		fmt.Println(part1(parse("1002,4,3,4,33"), 1))
		// fmt.Println(part1(parse("3,9,8,9,10,9,4,9,99,-1,8"), 5))

		test := parse("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99")
		// fmt.Println(part1(parse(), 12, 2, []int{8}))
		fmt.Println(part1(test, 1))
		fmt.Println(part1(test, 8))
		fmt.Println(part1(test, 9))

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

		fmt.Println(part1(prog, 1))
		fmt.Println(part1(prog, 5))
		// 7506531 too high

		// // noun, verb := 0, 0
		// found := make(chan int)
		// for noun := 0; noun < 100; noun++ {
		// 	for verb := 0; verb < 100; verb++ {
		// 		go func(noun, verb int) {
		// 			a := part1(prog, noun, verb)
		// 			if a == 19690720 {
		// 				found <- noun*100 + verb
		// 			}
		// 		}(noun, verb)
		// 	}
		// }
		// log.Println(<-found)
	}

}
