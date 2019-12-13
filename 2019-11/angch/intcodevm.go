package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var opcodes = map[int]string{
	1: "ADD", 2: "MUL", 3: "IN", 4: "OUT", 5: "JNE", 6: "JE", 7: "CMPL", 8: "CMPE", 9: "MOVBP", 99: "DIAG",
}

func dump(program map[int]int) {
	for i := 0; i < len(program); i++ {
		// FIXME
		fmt.Printf("%3d: %4d ", i, program[i])
		if i%10 == 9 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func IntCodeVM(prog map[int]int, input <-chan int, output chan<- int) int {
	program := make(map[int]int)
	for k, v := range prog {
		program[k] = v
	}
	lastOutput := 0
	ip := 0

	debug := false
	if debug {
		dump(program)
	}
	bp := 0

	for {
		if ip > len(program) || ip < 0 {
			log.Println("ss")
			return 0
		}
		opcode := program[ip]
		opcode, param1, param2, param3 := opcode%100, (opcode/100)%10, (opcode/1000)%10, (opcode/10000)%10

		a1, b1, d1 := program[ip+1], program[ip+2], program[ip+3]
		a0, b0, d0 := a1, b1, d1

		if debug {
			cparam1, cparam2 := "", ""
			if param1 == 0 {
				cparam1 = "*"
			} else if param1 == 2 {
				cparam1 = "bp+"
			}
			if param2 == 0 {
				cparam2 = "*"
			} else if param2 == 2 {
				cparam2 = "bp+"
			}
			log.Printf("%03d: %d %s %s%d %s%d %s%d", ip, program[ip], opcodes[opcode], cparam1, a0, cparam2, b0, "", d0)
		}

		switch param1 {
		case 0, 2:
			if param1 == 2 {
				a0 += bp
				a1 += bp
			}
			if a0 >= 0 {
				a0 = program[a0]
			} else {
				a0 = -999
			}
		}
		switch param2 {
		case 0, 2:
			if param2 == 2 {
				b0 += bp
				b1 += bp
			}
			if b0 >= 0 {
				b0 = program[b0]
			} else {
				b0 = -999
			}
		}
		if param3 == 2 {
			d0 += bp
			d1 += bp
		}

		switch opcode {
		case 1:
			program[d0], ip = a0+b0, ip+4
		case 2:
			program[d0], ip = a0*b0, ip+4
		case 3:
			program[a1], ip = <-input, ip+2
		case 4:
			ip = ip + 2
			lastOutput = a0
			output <- a0
		case 5:
			ip += 3
			if a0 != 0 {
				ip = b0
			}
		case 6:
			ip += 3
			if a0 == 0 {
				ip = b0
			}
		case 7:
			program[d1], ip = 0, ip+4
			if a0 < b0 {
				program[d1] = 1
			}
		case 8:
			program[d1], ip = 0, ip+4
			if a0 == b0 {
				program[d1] = 1
			}
		case 9:
			bp, ip = bp+a0, ip+2
		case 99:
			// log.Println("x x")
			close(output)
			return lastOutput
		default:
			return 1
			log.Fatal("fatal", program[ip])
		}
	}
}

func ParseIntCode(s string) map[int]int {
	s1 := strings.Split(s, ",")
	o := make(map[int]int, len(s1)+10)
	for k := range s1 {
		a, _ := strconv.Atoi(s1[k])
		o[k] = a
	}
	return o
}
