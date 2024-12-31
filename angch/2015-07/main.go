package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Wiring struct {
	op1  string
	op2  string
	op   string
	data uint16
}

var debug bool

func eval(wirings map[string]Wiring, key string, lvl int) uint16 {
	if key == "" {
		return 0
	}
	if key[0] >= '0' && key[0] <= '9' {
		var d uint16
		fmt.Sscanf(key, "%d", &d)
		return d
	}
	if debug {
		for i := 0; i < lvl; i++ {
			fmt.Print(" ")
		}
		fmt.Println("eval", key, wirings[key])
	}
	if wirings[key].op == "data" {
		return wirings[key].data
	}
	d1 := eval(wirings, wirings[key].op1, lvl+1)
	d2 := eval(wirings, wirings[key].op2, lvl+1)
	d := uint16(0)
	switch wirings[key].op {
	case "AND":
		d = d1 & d2
	case "OR":
		d = d1 | d2
	case "LSHIFT":
		d = d1 << d2
	case "RSHIFT":
		d = d1 >> d2
	case "NOT":
		d = ^d1
	case "":
		d = d1
	}
	wirings[key] = Wiring{
		data: d,
		op:   "data",
	}
	if debug {
		for i := 0; i < lvl; i++ {
			fmt.Print(" ")
		}

		fmt.Println("evald", key, wirings[key])
	}
	return d
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	wirings := make(map[string]Wiring)
	for scanner.Scan() {
		line := scanner.Text()
		left := ""
		var dest string
		ops := strings.Split(line, " -> ")
		left = ops[0]
		dest = ops[1]
		// fmt.Println(left, right)
		num, err := strconv.ParseUint(left, 10, 16)
		if err == nil {
			// 123 -> x
			wirings[dest] = Wiring{
				data: uint16(num),
				op:   "data",
			}
		} else {
			ops2 := strings.Split(left, " ")
			// fmt.Println(" ops2", strings.Join(ops2, ","))
			w := Wiring{
				op1: ops2[0],
				op2: "",
				op:  "",
			}
			if len(ops2) > 1 {
				w.op = ops2[1]
			}
			if len(ops2) == 3 {
				w.op2 = ops2[2]
			} else {
				if ops2[0] == "NOT" {
					w.op1 = ops2[1]
					w.op = "NOT"
				}
			}
			wirings[dest] = w
		}
	}
	if debug {
		fmt.Println(wirings)
	}

	newwire := make(map[string]Wiring)
	for k, v := range wirings {
		newwire[k] = v
	}

	part1 := eval(wirings, "a", 0)
	newwire["b"] = Wiring{
		data: part1,
		op:   "data",
	}
	part2 := eval(newwire, "a", 0)
	fmt.Println(part1, part2)
}
