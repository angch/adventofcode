package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func do(fileName string) (int, int) {
	ret1, ret2 := 0, 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	memory := make(map[int]uint64, 100)
	mask1 := 0
	mask2 := 0
	for scanner.Scan() {
		l := scanner.Text()

		tok := strings.Split(l, " ")
		mask := ""
		sp := 0
		dst := 0

		if len(tok) < 2 {
			continue
		}
		if tok[0] == "mask" {
			mask = tok[2]
			// log.Println("mask", mask)
			mask1 = 0
			mask2 = 0
			for _, v := range mask {
				switch v {
				case 'X':
					mask1 <<= 1
					mask2 <<= 1
					mask2 |= 1
				case '1':
					mask1 <<= 1
					mask2 <<= 1
					mask1 |= 1
				case '0':
					mask1 <<= 1
					mask2 <<= 1
				}
			}
			// log.Println("masks ", mask1, mask2)
		} else {
			fmt.Sscanf(tok[0], "mem[%d]", &sp)
			fmt.Sscanf(tok[2], "%d", &dst)
			// log.Println(sp, dst)
			val := (dst & mask2) | mask1
			memory[sp] = uint64(val)
			// log.Println(val)
		}
		// fmt.Sscanf(l, "%s %d", &opcode.Op, &opcode.Arg)
	}
	// log.Println(memory)

	for _, v := range memory {
		ret1 += int(v)
	}
	return ret1, ret2
}

func do2(fileName string) (int, int) {
	ret1, ret2 := 0, 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	memory := make(map[int]uint64, 100)
	masks1 := make([]int, 0)
	masks2 := make([]int, 0)
	// masks3 := make([]int, 0)
	for scanner.Scan() {
		l := scanner.Text()

		tok := strings.Split(l, " ")
		mask := ""
		sp := 0
		dst := 0

		if len(tok) < 2 {
			continue
		}
		if tok[0] == "mask" {
			mask = tok[2]
			// log.Println("mask", mask)
			masks1 = make([]int, 1)
			masks2 = make([]int, 1)
			// masks3 = make([]int, 1)

			for k, v := range mask {
				switch v {
				case 'X':
					for k := range masks1 {
						masks1[k] <<= 1
						masks2[k] <<= 1
						// masks2[k] |= 1
					}
					masks1 = append(masks1, masks1...)
					masks2 = append(masks2, masks2...)

					for k = len(masks1) / 2; k < len(masks1); k++ {
						masks1[k] |= 1
						// masks2[k] &= 1
					}
					// for k = 0; k < len(masks1)/2; k++ {
					// 	// masks2[k] |= 1
					// }
				case '1':
					for k := range masks1 {
						masks1[k] <<= 1
						masks2[k] <<= 1
						masks1[k] |= 1
					}
				case '0':
					for k := range masks1 {
						masks1[k] <<= 1
						masks2[k] <<= 1
						masks2[k] |= 1
					}
				}
			}
			// for k := range masks1 {
			// 	log.Printf("mask %b %b\n", masks1[k]+1024, masks2[k]+1024)
			// }
			// log.Println("masks ", masks1, masks2)
		} else {
			fmt.Sscanf(tok[0], "mem[%d]", &sp)
			fmt.Sscanf(tok[2], "%d", &dst)
			// log.Println(sp, dst)
			// val := (sp & mask2) | mask1
			val := dst

			// log.Println("sp is ", sp)
			// oldsp :=
			for k := range masks1 {
				sp2 := (sp & masks2[k]) | masks1[k]
				// log.Printf("sp %d %b %b\n", sp2, masks2[k], masks1[k])
				memory[sp2] = uint64(val)
			}
			// log.Println(val)
		}
		// fmt.Sscanf(l, "%s %d", &opcode.Op, &opcode.Arg)
	}
	// log.Println(memory)

	for _, v := range memory {
		ret1 += int(v)
	}
	return ret1, ret2
}

func main() {
	log.Println(do("test.txt"))
	log.Println(do2("test2.txt"))
	// log.Println(do("test2.txt"))
	log.Println(do("input.txt"))
	log.Println(do2("input.txt"))
}
