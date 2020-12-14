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

	memory := make(map[int]uint64)
	mask1, mask2 := 0, 0
	for scanner.Scan() {
		l := scanner.Text()

		tok := strings.Split(l, " ")
		mask, sp, dst := "", 0, 0

		if tok[0] == "mask" {
			mask, mask1, mask2 = tok[2], 0, 0
			// log.Println("mask", mask)
			for _, v := range mask {
				mask1 <<= 1
				mask2 <<= 1
				switch v {
				case 'X':
					mask2 |= 1
				case '1':
					mask1 |= 1
				}
			}
		} else {
			fmt.Sscanf(tok[0], "mem[%d]", &sp)
			fmt.Sscanf(tok[2], "%d", &dst)
			val := (dst & mask2) | mask1
			memory[sp] = uint64(val)
		}
	}

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

	memory := make(map[int]uint64)
	masks1 := make([]int, 0)
	masks2 := make([]int, 0)
	// masks3 := make([]int, 0)
	for scanner.Scan() {
		l := scanner.Text()

		tok := strings.Split(l, " ")
		mask, sp, dst := "", 0, 0
		if tok[0] == "mask" {
			mask = tok[2]
			masks1 = make([]int, 1)
			masks2 = make([]int, 1)

			for _, v := range mask {
				for k := range masks1 {
					masks1[k] <<= 1
					masks2[k] <<= 1
				}
				switch v {
				case 'X':
					masks1 = append(masks1, masks1...)
					masks2 = append(masks2, masks2...)

					for k := 0; k < len(masks1)/2; k++ {
						masks1[k] |= 1
					}
				case '1':
					for k := range masks1 {
						masks1[k] |= 1
					}
				case '0':
					for k := range masks1 {
						masks2[k] |= 1
					}
				}
			}
		} else {
			fmt.Sscanf(tok[0], "mem[%d]", &sp)
			fmt.Sscanf(tok[2], "%d", &dst)

			for k := range masks1 {
				sp2 := (sp & masks2[k]) | masks1[k]
				memory[sp2] = uint64(dst)
			}
		}
	}

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
