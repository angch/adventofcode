package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var verbose = false

func day7(file string) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	curdir, cmd := "", ""
	dirsize := make(map[string]int)
	for scanner.Scan() {
		t := scanner.Text()

		words := strings.Split(t, " ")
		if words[0] == "$" {
			switch words[1] {
			case "cd":
				if strings.HasPrefix(words[2], "/") {
					curdir = words[2]
				} else {
					if words[2] == ".." {
						a := strings.Split(curdir, "/")
						a = a[:len(a)-1]
						curdir = strings.Join(a, "/")
					} else {
						curdir += "/" + words[2]
					}
				}
				if curdir == "/" {
					curdir = ""
				}
				dirsize[curdir] += 0
				cmd = ""
			case "ls":
				cmd = "ls"
			}
		} else {
			switch cmd {
			case "ls":
				if words[0] != "dir" {
					size, _ := strconv.Atoi(words[0])
					dirs := strings.Split(curdir, "/")
					for k := 0; k < len(dirs); k++ {
						dirsize[strings.Join(dirs[0:k+1], "/")] += size
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
	sizes := make([]int, 0)
	for _, v := range dirsize {
		if v < 100000 {
			part1 += v
		}
		sizes = append(sizes, v)
	}
	sort.Ints(sizes)
	disk := 30000000 - (70000000 - dirsize[""])
	if verbose {
		log.Println("looking for disk", disk)
	}
	for _, v := range sizes {
		if v > disk {
			part2 = v
			break
		}
	}
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
