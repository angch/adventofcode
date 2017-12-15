package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	id    int
	links map[int]bool
	color int
}

func do(inputs []string) (int, int) {
	nodes := make(map[int]node)
	for _, line := range inputs {
		id := 0
		tokens := strings.Split(line, " <-> ")
		fmt.Sscanf(tokens[0], "%d", &id)
		linksString := strings.Split(tokens[1], ",")
		links := make(map[int]bool, 0)
		for _, linkString := range linksString {
			ls := strings.Trim(linkString, " ") // #$%^& STUPID BUG!
			link, _ := strconv.Atoi(ls)
			links[link] = true
		}

		nodes[id] = node{
			id:    id,
			links: links,
		}
		//fmt.Println(nodes, linksString)
	}

	if true {
		for id, n := range nodes {
			for l, _ := range n.links {
				nodes[l].links[id] = true
			}
		}
	}

	c := color(nodes, 0, 1)

	nextColor := 0
	count := 0
a:
	for {
		count = 0
		for id, n := range nodes {
			if n.color == 1 {
				count++
			} else if n.color == 0 {
				nextColor++
				color(nodes, id, nextColor)
				continue a
			}
		}
		break
	}

	if false {
		// Debugging stuff
		fmt.Println("check", count)
		keys := make([]int, 0)
		for id, _ := range nodes {
			keys = append(keys, id)
		}
		sort.Ints(keys)
		for key, _ := range keys {
			n := nodes[key]
			if n.color == 1 {
				n.color = 9999
			} else if n.links[key] {
				n.color = 8888
			}
			fmt.Println("nn", n.id, n.color, n.links)
		}
	}
	return c, nextColor
}

func color(nodes map[int]node, id, c int) int {
	count := 0
	n := nodes[id]
	if n.color != c {
		n.color = c
		count++
		//fmt.Println("color", id, c)
	}
	n.color = c
	nodes[id] = n

	for l, _ := range nodes[id].links {
		if nodes[l].color == 0 {
			count += color(nodes, l, c)
		} else if nodes[l].color != c {
			log.Fatal("Not supposed to happen", l)
		}
	}
	return count
}

func main() {
	if false {
		file, _ := ioutil.ReadFile("test.txt")
		a, b := do(strings.Split(string(file), "\n"))
		fmt.Println(a, b)
	}

	if true {
		file, _ := ioutil.ReadFile("input.txt")
		a, b := do(strings.Split(string(file), "\n"))
		fmt.Println(a, b)
	}
}
