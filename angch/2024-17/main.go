package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type VM struct {
	regA    int
	regB    int
	regC    int
	program []int
	ip      int
}

func (vm *VM) step() int {
	if vm.ip+1 >= len(vm.program) {
		log.Println("Halt")
		return -2
	}
	opcode := vm.program[vm.ip]
	operand := vm.program[vm.ip+1]
	// fmt.Println("IP", ip, "Opcode", opcode, "Operand", operand)
	// fmt.Println("Registers", registers)

	combo := 0
	switch operand {
	case 0, 1, 2, 3:
		combo = operand
	case 4:
		combo = vm.regA
	case 5:
		combo = vm.regB
	case 6:
		combo = vm.regC
	}
	switch opcode {
	case 0: // adv
		num := vm.regA
		vm.regA = num >> uint(combo)
	case 1: // bxl
		vm.regB ^= operand
	case 2: // bst
		vm.regB = combo & 7
	case 3: // jnz
		a := vm.regA
		if a != 0 {
			vm.ip = operand
			return -1
		}
	case 4: // bxc
		vm.regB ^= vm.regC
	case 5: // out
		vm.ip += 2
		return combo & 7
	case 6: // bdv
		num := vm.regA
		vm.regB = num >> uint(combo)
	case 7: // cdv
		num := vm.regA
		vm.regC = num >> uint(combo)
	}
	vm.ip += 2
	return -1
}

func day17(file string) (part1 string, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	registers := map[byte]int{}
	program := []int{}

	for scanner.Scan() {
		t := scanner.Text()
		_ = t
		if t == "" {
			break
		}
		words := strings.Split(t, " ")
		r := words[1][0]
		v, _ := strconv.Atoi(words[2])
		registers[r] = v

	}
	scanner.Scan()
	t := scanner.Text()
	_, r, _ := strings.Cut(t, " ")
	for _, v := range strings.Split(r, ",") {
		v1, _ := strconv.Atoi(v)
		program = append(program, v1)
	}

	log.Println(registers)
	log.Println(program)
	_ = program

	ip := 0

	vm := VM{regA: registers['A'], regB: registers['B'], regC: registers['C'], program: program, ip: ip}
	output := []int{}
	for {
		out := vm.step()
		if out == -2 {
			break
		}
		if out >= 0 {
			output = append(output, out)
		}
	}

	matched := 0
	t2 := time.Now()
	maxmatched := 0
a:
	for {
		vm.regA = part2
		vm.regB = registers['B']
		vm.regC = registers['C']
		vm.ip = 0

		if time.Since(t2) > 5*time.Second {
			fmt.Println("Part2", part2, maxmatched)
			t2 = time.Now()
		}

		expect := program[matched]
		// 2751
		outcount := 0
		for {
			out := vm.step()
			if out == -2 {
				part2 += 1
				continue a
			}
			if out == -1 {
				continue
			}
			if out >= 0 {
				outcount++
				if outcount >= len(program) {
					log.Println("Part2", part2)
					break a
				}
				continue
			}
			if expect == out {
				matched++
				if matched > maxmatched {
					maxmatched = matched
					fmt.Println("Max matched", maxmatched, "Part2", part2)
				}
				if matched == len(program) {
					break a
				}
				expect = program[matched]
			} else {
				matched = 0
				part2 += 2751
				continue a
			}
		}
	}

	out2 := []string{}
	for _, v := range output {
		out2 = append(out2, fmt.Sprintf("%d", v))
	}
	part1 = strings.Join(out2, ",")
	// log.Println("Final registers", registers)
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	// part1, part2 := day17("test.txt")
	// fmt.Println(part1, part2)
	// if part1 != "0,1,2" || part2 != "" {
	// 	log.Fatal("Test failed ", part1, part2)
	// }
	// part1, part2 = day17("test1.txt")
	// fmt.Println(part1, part2)
	// if part1 != "4,2,5,6,7,7,7,7,3,1,0" || part2 != "" {
	// 	log.Fatal("Test failed ", part1, part2)
	// }
	// part1, part2 = day17("test0.txt")
	// fmt.Println(part1, part2)
	// if part1 != "4,6,3,5,6,3,5,2,1,0" || part2 != "" {
	// 	log.Fatal("Test failed ", part1, part2)
	// }
	// part1, part2 := day17("test3.txt")
	// fmt.Println(part1, part2)
	// if part1 != "0,1,2" || part2 != "" {
	// 	log.Fatal("Test failed ", part1, part2)
	// }
	// Not: 2,4,0,4,6,2,3,3,4
	// Not: 34 11142549
	fmt.Println(day17("input.txt"))
	fmt.Println("Elapsed time:", time.Since(t1))
}
