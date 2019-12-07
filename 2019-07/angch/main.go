package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var opcodes = map[int]string{
	1: "ADD", 2: "MUL", 3: "IN", 4: "OUT", 5: "JNE", 6: "JE", 7: "CMPL", 8: "CMPE", 99: "DIAG",
}

func dump(program []int) {
	for i := 0; i < len(program); i++ {
		fmt.Printf("%3d: %4d ", i, program[i])
		if i%10 == 9 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func part1(prog []int, input <-chan int, output chan<- int) int {
	program := make([]int, len(prog))
	copy(program, prog)
	lastOutput := 0
	ip := 0

	debug := false
	if debug {
		dump(program)
	}

	for {
		if ip > len(program) || ip < 0 {
			log.Println("ss")
			return 0
		}
		opcode := program[ip]
		opcode, param1, param2 := opcode%100, (opcode/100)%10, (opcode/1000)%10

		a1, b1, d1 := program[ip+1], program[ip+2], program[ip+3]
		a0, b0, d0 := a1, b1, d1

		if debug {
			cparam1, cparam2 := "", ""
			if param1 == 0 {
				cparam1 = "*"
			}
			if param2 == 0 {
				cparam2 = "*"
			}
			log.Printf("%03d: %s %s%d %s%d %s%d", ip, opcodes[opcode], cparam1, a0, cparam2, b0, "", d0)
		}

		if param1 == 0 {
			if a0 < len(program) && a0 >= 0 {
				a0 = program[a0]
			} else {
				a0 = -999
			}
		}
		if param2 == 0 {
			if b0 < len(program) && b0 >= 0 {
				b0 = program[b0]
			} else {
				b0 = -999
			}
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
		case 99:
			// log.Println("x x")
			return lastOutput
		default:
			return 1
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

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func part0(program []int, part int) int {
	var signals [][]int
	if part == 1 {
		signals = permutations([]int{0, 1, 2, 3, 4})
	} else {
		signals = permutations([]int{5, 6, 7, 8, 9})
	}

	result := make(chan int)
	for _, signal := range signals {
		go func(signal []int) {
			signalchan := make([](chan int), 6)
			for i := range signalchan {
				signalchan[i] = make(chan int, 1)
			}

			wg := sync.WaitGroup{}
			wg.Add(5)
			go func() { part1(program, signalchan[0], signalchan[1]); wg.Done() }()
			go func() { part1(program, signalchan[1], signalchan[2]); wg.Done() }()
			go func() { part1(program, signalchan[2], signalchan[3]); wg.Done() }()
			go func() { part1(program, signalchan[3], signalchan[4]); wg.Done() }()
			go func() { result <- part1(program, signalchan[4], signalchan[0]); wg.Done() }()

			for i := range signal {
				signalchan[i] <- signal[i]
			}
			signalchan[0] <- 0

			wg.Wait()

			// result <- (<-signalchan[0])
		}(signal)
	}
	max := 0
	for range signals {
		r := <-result
		// log.Println(r)
		if r > max {
			max = r
		}
	}
	log.Println("max", max)
	return max
}

func main() {
	if false {
		// Part 1
		program := parse("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0")
		part0(program, 1)

		program = parse("3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0")
		part0(program, 1)

		program = parse("3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0")
		part0(program, 1)
	}
	if true {
		// Part 2
		program := parse("3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5")
		part0(program, 2)

		program = parse("3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10")
		part0(program, 2)
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
		fmt.Println("Part 1:")
		part0(prog, 1)
		fmt.Println("Part 2:")
		part0(prog, 2)

		// fmt.Println("Part 1", part1(prog, 1))
		// fmt.Println("Part 2", part1(prog, 5))
	}
}
