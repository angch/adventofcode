package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

type Opcode struct {
	Op   string
	Arg1 string
	Arg2 string
}

func runner1(instr []Opcode) int {
	ip := 0
	registers := make(map[string]int)
	playedFreq := 0
	ret := 0
	// log.Println(instr)
	// return 0, 0
a:
	for ip < len(instr) {
		opcode := instr[ip]

		arg2, err := strconv.Atoi(opcode.Arg2)
		if err != nil {
			arg2 = registers[opcode.Arg2]
		}
		arg1, err := strconv.Atoi(opcode.Arg1)
		if err != nil {
			arg1 = registers[opcode.Arg1]
		}
		switch opcode.Op {
		case "set":
			registers[opcode.Arg1] = arg2
		case "add":
			registers[opcode.Arg1] += arg2
		case "mul":
			registers[opcode.Arg1] *= arg2
		case "mod":
			registers[opcode.Arg1] = arg1 % arg2
		case "snd":
			playedFreq = arg1
		case "rcv":
			if registers[opcode.Arg1] != 0 {
				ret = playedFreq
				break a
			}
		case "jgz":
			if arg1 > 0 {
				ip += arg2
				continue a
			}
		}
		ip++
	}
	return ret

}

func runner2(instr []Opcode, input <-chan int, output chan<- int, done chan bool, p int, counter *int64, wait *int64) int {
	ip := 0
	registers := make(map[string]int)
	registers["p"] = p
	// playedFreq := 0
	ret := 0
	// log.Println(instr)
	// return 0, 0
a:
	for ip < len(instr) {
		opcode := instr[ip]

		arg2, err := strconv.Atoi(opcode.Arg2)
		if err != nil {
			arg2 = registers[opcode.Arg2]
		}
		arg1, err := strconv.Atoi(opcode.Arg1)
		if err != nil {
			arg1 = registers[opcode.Arg1]
		}
		// if p == 1 {
		// 	// 	// log.Println(p, opcode, opcode.Arg1, registers)
		// 	log.Println(p, registers)
		// }
		switch opcode.Op {
		case "set":
			registers[opcode.Arg1] = arg2
		case "add":
			registers[opcode.Arg1] += arg2
		case "mul":
			registers[opcode.Arg1] *= arg2
		case "mod":
			registers[opcode.Arg1] = arg1 % arg2
		case "snd":
			if counter != nil {
				atomic.AddInt64(counter, 1)
			}
			output <- arg1
			// log.Println(p, opcode, opcode.Arg1, arg1, registers)
			// if p == 0 {
			// 	log.Println("snd", arg1)
			// }

		case "rcv":
			atomic.AddInt64(wait, 1)
		b:
			for {
				select {
				case a := <-input:
					// log.Println(p, opcode, a, opcode.Arg1, registers)
					registers[opcode.Arg1] = a
					// log.Println(p, opcode, a, opcode.Arg1, registers)
					atomic.AddInt64(wait, -1)
					// if p == 1 {
					// 	log.Println("recv", a, registers)
					// }
					break b
				case <-time.After(100 * time.Millisecond):
					w := atomic.LoadInt64(wait)
					if w >= 2 {
						if counter != nil {
							// log.Println("Deadlock", *counter)
							// done <- true
							// return ret
						}
						done <- true
						// done <- true
						return ret
					}
					// log.Println(p, "wait is", *wait)
				}
			}
		case "jgz":
			if arg1 > 0 {
				ip += arg2
				continue a
			}
		}
		ip++
	}
	done <- true
	return ret
}

func do(fileName string) (int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	instr := make([]Opcode, 0)
	// ip := 0
	for scanner.Scan() {
		l := scanner.Text()
		opcode := Opcode{}

		fmt.Sscanf(l, "%s %s %s", &opcode.Op, &opcode.Arg1, &opcode.Arg2)
		instr = append(instr, opcode)
	}
	// log.Println(instr)

	ret1 := runner1(instr)

	input := make(chan int, 10000)
	output := make(chan int, 10000)
	done := make(chan bool, 2)

	counter := int64(0)
	wait := int64(0)
	go runner2(instr, input, output, done, 0, nil, &wait)
	go runner2(instr, output, input, done, 1, &counter, &wait)

	oldcounter := int64(0)
	donecounter := 0
a:
	for {
		select {
		case <-done:
			donecounter++
			if donecounter >= 2 {
				return ret1, int(counter)
			}

		case <-time.After(5 * time.Second):
			if oldcounter == counter {
				break a
			}
			// log.Println("x", len(input), len(output), counter, oldcounter)
			oldcounter = counter
		}
	}
	//80235 too high
	return ret1, int(counter)
}

func main() {
	log.Println(do("test.txt"))
	log.Println(do("test2.txt"))
	log.Println(do("input.txt"))
}
