package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Instruction struct {
	reg string
	cmd string
	num int

	reg1 string
	cnd  string
	prm  int
}

func main() {
	//log.SetOutput(ioutil.Discard)

	if true {
		test1, test2 := advent08a("test1.txt")
		fmt.Println(test1, test2)
		if test1 != 1 {
			log.Fatal("err")
		}
		o, o2 := advent08a("input.txt")
		fmt.Println(o, o2)
	}
}

func advent08a(fileName string) (int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	instructions := make([]Instruction, 0)
	for scanner.Scan() {
		t := scanner.Text()

		i := Instruction{}
		n := ""
		fmt.Sscanf(t, "%s %s %d %s %s %s %d", &i.reg, &i.cmd, &i.num, &n, &i.reg1, &i.cnd, &i.prm)
		instructions = append(instructions, i)
	}
	reg := make(map[string]int)

	max2 := -1
	for _, i := range instructions {
		reg1 := reg[i.reg1]
		t := false
		switch i.cnd {
		case "<":
			t = reg1 < i.prm
		case ">":
			t = reg1 > i.prm
		case ">=":
			t = reg1 >= i.prm
		case "<=":
			t = reg1 <= i.prm
		case "==":
			t = reg1 == i.prm
		case "!=":
			t = reg1 != i.prm
		default:
			log.Fatal(i.cnd)
		}

		if t {
			if i.cmd == "inc" {
				reg[i.reg] += i.num
			} else {
				reg[i.reg] -= i.num
			}
			if reg[i.reg] > max2 {
				max2 = reg[i.reg]
			}
		}
	}

	max := -1
	for _, r := range reg {
		if r > max {
			max = r
		}
	}
	return max, max2
}
