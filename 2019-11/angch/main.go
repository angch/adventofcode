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

type Coord struct {
	x, y int
}

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

func AmpRunner(program map[int]int, part int) int {
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

			// Wire up all the "Amp"s and connect the inputs together.
			// Note all of them will start running until blocked and waiting
			// on their inputs (phase first):
			go func() { IntCodeVM(program, signalchan[0], signalchan[1]); wg.Done() }()
			go func() { IntCodeVM(program, signalchan[1], signalchan[2]); wg.Done() }()
			go func() { IntCodeVM(program, signalchan[2], signalchan[3]); wg.Done() }()
			go func() { IntCodeVM(program, signalchan[3], signalchan[4]); wg.Done() }()
			go func() { result <- IntCodeVM(program, signalchan[4], signalchan[0]); wg.Done() }()

			// Load all phases into the inputs to the Amps
			for i := range signal {
				signalchan[i] <- signal[i]
			}
			// Signal everyone to start processing
			signalchan[0] <- 0
			wg.Wait()
		}(signal)
	}
	max := 0
	for range signals {
		r := <-result
		if r > max {
			max = r
		}
	}
	log.Println("max", max)
	return max
}

func paint(prog map[int]int, init int) {
	output := make(chan int)
	input := make(chan int, 1)
	//input <- 1

	go IntCodeVM(prog, input, output)

	grid := make(map[Coord]int)
	curloc := Coord{0, 0}
	curdir := Coord{0, -1} // Up

	globalpainted := 0

	minloc, maxloc := Coord{0, 0}, Coord{0, 0}

	if init > 0 {
		grid[curloc] = init
	}
	for {
		paintold, painted := grid[curloc]
		input <- paintold
		paintnew, more := <-output
		if !more {
			break
		}

		grid[curloc] = paintnew
		if !painted {
			globalpainted++
		}

		turn := <-output
		if turn == 0 {
			curdir.x, curdir.y = curdir.y, -curdir.x
		} else {
			curdir.x, curdir.y = -curdir.y, curdir.x
		}
		curloc.x, curloc.y = curloc.x+curdir.x, curloc.y+curdir.y
		if minloc.x > curloc.x {
			minloc.x = curloc.x
		}
		if minloc.y > curloc.y {
			minloc.y = curloc.y
		}
		if maxloc.x < curloc.x {
			maxloc.x = curloc.x
		}
		if maxloc.y < curloc.y {
			maxloc.y = curloc.y
		}
	}
	log.Println("Count", globalpainted, minloc, maxloc) // 1236 too low

	for y := minloc.y; y <= maxloc.y; y++ {
		for x := minloc.x; x <= maxloc.x; x++ {
			if grid[Coord{x, y}] == 1 {
				fmt.Print("##")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

}

func main() {
	if true {
		fileName := "input.txt"

		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Scan()
		prog := ParseIntCode(scanner.Text())

		fmt.Println("Part 1:")
		paint(prog, 0)
		fmt.Println("Part 2:")
		paint(prog, 1)

	}
}
