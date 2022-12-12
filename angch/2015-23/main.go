package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instr struct {
	op     string
	param1 string
	param2 int
}

func sim(instructions []instr, registers map[string]int) int {
	for ip := 0; ip < len(instructions); ip++ {
		op := instructions[ip]
		// fmt.Println(ip, op, registers)
		switch op.op {
		case "hlf":
			registers[op.param1] /= 2
		case "tpl":
			registers[op.param1] *= 3
		case "inc":
			registers[op.param1]++
		case "jmp":
			ip += op.param2 - 1
		case "jie":
			if registers[op.param1]%2 == 0 {
				// fmt.Println("jumping")
				ip += op.param2 - 1
			}
		case "jio":
			// fmt.Println("j", op.param1, registers[op.param1], registers)
			if registers[op.param1] == 1 {
				// fmt.Println("jumping1")
				ip += op.param2 - 1
			}
		}
	}
	return registers["b"]
}
func day22(file string) (int, int) {
	part1, part2 := 0, 0
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	instructions := []instr{}
	for scanner.Scan() {
		t := strings.ReplaceAll(scanner.Text(), ",", "")
		words := strings.Split(t, " ")
		switch words[0] {
		case "inc", "tpl", "hlf":
			instructions = append(instructions, instr{op: words[0], param1: words[1]})
		case "jmp":
			p, _ := strconv.Atoi(words[1])
			instructions = append(instructions, instr{op: words[0], param2: p})
		case "jie", "jio":
			p, _ := strconv.Atoi(words[2])
			instructions = append(instructions, instr{op: words[0], param1: words[1], param2: p})
		}
	}
	fmt.Printf("%+v\n", instructions)

	registers := map[string]int{"a": 0, "b": 0}
	part1 = sim(instructions, registers)
	registers = map[string]int{"a": 1, "b": 0}
	part2 = sim(instructions, registers)

	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	part1, part2 := day22("test.txt")
	fmt.Println(part1, part2)
	fmt.Println(day22("input.txt"))
}
