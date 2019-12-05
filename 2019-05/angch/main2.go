package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var opcodes = map[int]string{
	1:  "ADD",
	2:  "MUL",
	3:  "IN",
	4:  "OUT",
	5:  "JNE",
	6:  "JE",
	7:  "CMPL",
	8:  "CMPE",
	99: "DIAG",
}

func part1(prog []int, input int) int {
	program := make([]int, len(prog))
	copy(program, prog)
	ip := 0
	output := make([]int, 0)

	debug := false

	if debug {
		for i := 0; i < len(program); i++ {
			fmt.Printf("%3d: %4d ", i, program[i])
			if i%10 == 9 {
				fmt.Println()
			}
		}
		fmt.Println()
	}

	for {
		opcode := program[ip] % 100
		code := program[ip] / 100

		param1 := code % 10
		param2 := (code / 10) % 10
		param3 := (code / 100) % 10
		if param3 > 0 {
			log.Println("param3 in use!")
		}
		a, b, d := program[ip+1], program[ip+2], program[ip+3]
		cparam1, cparam2 := "", ""
		if param1 == 0 {
			cparam1 = "*"
		}
		if param2 == 0 {
			cparam2 = "*"
		}
		if debug {
			log.Printf("%03d: %s %s%d %s%d %s%d", ip, opcodes[opcode], cparam1, a, cparam2, b, "", d)
		}

		if param1 == 0 {
			if a < len(program) {
				a = program[a]
			} else {
				a = -999
			}
		}
		if param2 == 0 {
			if b < len(program) {
				b = program[b]
			} else {
				b = -999
			}
		}

		switch opcode {
		case 1:
			program[d] = a + b
			if debug {
				log.Println(ip, "ADD", a, b, d, program[d])
			}
			ip += 4
		case 2:
			program[d] = a * b
			if debug {
				log.Println(ip, "MUL", a, b, d, program[d])
			}
			ip += 4
		case 3:
			// program[a] = input
			program[program[ip+1]] = input
			if debug {
				log.Println(ip, "IN", a, program[a])
			}
			ip += 2
		case 4:
			output = append(output, a)
			if debug {
				log.Println(ip, "OUT", a, b, d)
			}
			ip += 2
		case 5:
			oip := ip
			if a != 0 {
				ip = b
			} else {
				ip += 3
			}
			if debug {
				log.Println(oip, "JNE", a, b, d, ip)
			}
		case 6:
			oip := ip
			if a == 0 {
				ip = b
			} else {
				ip += 3
			}
			if debug {
				log.Println(oip, "JE", a, b, d)
			}

		case 7:
			if a < b {
				program[d] = 1
			} else {
				program[d] = 0
			}
			if debug {
				log.Println(ip, "CMPL", a, b, d, program[d])
			}

			ip += 4
		case 8:
			if a == b {
				program[d] = 1
			} else {
				program[d] = 0
			}
			if debug {
				log.Println(ip, "CMPE", a, b, d, program[d])
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
	o := make([]int, len(s1)+10)
	for k := range s1 {
		a, _ := strconv.Atoi(s1[k])
		o[k] = a
	}
	return o
}

func main() {
	if true {
		// Part 1
		fmt.Println(part1(parse("1002,4,3,4,33"), 1))

		// fmt.Println(part1(parse("3,9,8,9,10,9,4,9,99,-1,8"), 5))

		// Part 2
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

		fmt.Println("Part 1", part1(prog, 1))
		fmt.Println("Part 2", part1(prog, 5))
	}

}
