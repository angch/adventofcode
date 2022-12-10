package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var verbose = false

func day7(file string) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	curdir, cmd := []string{}, ""
	dirsize := make(map[string]int)
	for scanner.Scan() {
		t := scanner.Text()

		words := strings.Split(t, " ")
		if words[0] == "$" {
			switch words[1] {
			case "cd":
				if strings.HasPrefix(words[2], "/") {
					curdir = strings.Split(words[2], "/")[1:]
				} else {
					if words[2] == ".." {
						curdir = curdir[:len(curdir)-1]
					} else {
						curdir = append(curdir, words[2])
					}
				}
				cmd = ""
			case "ls":
				cmd = "ls"
			}
		} else {
			switch cmd {
			case "ls":
				if words[0] != "dir" {
					size, _ := strconv.Atoi(words[0])
					for k := range curdir {
						dirsize[strings.Join(curdir[:k+1], "/")] += size
					}
				}
			}
		}
		if verbose {
			fmt.Println("After ", t)
			fmt.Println("curdir[", curdir, "] curcmd [", cmd, "]")
		}
	}
	if verbose {
		fmt.Println(dirsize)
	}
	bestSize := dirsize[""]
	disk := 30000000 - (70000000 - bestSize)
	if verbose {
		log.Println("looking for disk", disk)
	}
	for _, v := range dirsize {
		if v < 100000 {
			part1 += v
		}
		if v > disk && v < bestSize {
			bestSize = v
		}
	}
	part2 = bestSize
	return part1, part2
}

func main() {
	part1, part2 := day7("test.txt")
	fmt.Println(part1, part2)
	if part1 != 95437 || part2 != 24933642 {
		log.Fatal("Test failed")
	}
	fmt.Println(day7("input.txt"))
}
