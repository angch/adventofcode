package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Opcode struct {
	Op  string
	Arg int
}

func runner(vm []Opcode) (int, bool) {
	acc, ip := 0, 0
	history := make(map[int]bool)
	for {
		if history[ip] {
			return acc, false
		}
		history[ip] = true
		if ip >= len(vm) {
			return acc, true
		}
		opcode := vm[ip]
		switch opcode.Op {
		case "acc":
			acc += opcode.Arg
		case "jmp":
			ip += opcode.Arg - 1
		}
		ip++
	}
}

func do(fileName string) (int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	vm := make([]Opcode, 0)
	for scanner.Scan() {
		l := scanner.Text()
		opcode := Opcode{}
		fmt.Sscanf(l, "%s %d", &opcode.Op, &opcode.Arg)
		vm = append(vm, opcode)
	}
	ret1, _ := runner(vm)
	ret2 := make(chan int)

	for k := 0; k < len(vm); k++ {
		if vm[k].Op == "jmp" {
			go func(k int) {
				vm2 := make([]Opcode, len(vm))
				copy(vm2, vm)
				vm2[k].Op = "nop"

				acc, term := runner(vm2)
				if term {
					ret2 <- acc
				}
			}(k)
		}
	}

	return ret1, <-ret2
}

func main() {
	// log.Println(do("test.txt"))
	// log.Println(do("test2.txt"))
	log.Println(do("input.txt"))
}
