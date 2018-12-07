package main

// WTF ugly

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

// MOUBACDEFRGHIJKNPQSTXVWYZL
// MOUBNYITKXZFHQRJDASGCPEVWL
// MNOUBYITKXZFHQRJDASGCPEVWL

type Node struct {
	From string
	To   string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile((`^Step\s+(\w+) must be finished before step (\w+)`))

	nodes := make([]Node, 0)
	prereq := make(map[string][]string)
	availMap := make(map[string]bool)
	dot := false
	for scanner.Scan() {
		v := scanner.Text()
		a := re.FindAllStringSubmatch(v, -1)
		n := Node{a[0][1], a[0][2]}
		nodes = append(nodes, n)
		if prereq[n.To] == nil {
			prereq[n.To] = make([]string, 0)
		}
		prereq[n.To] = append(prereq[n.To], n.From)
		if dot {
			fmt.Println(n.From, "->", n.To)
		}
		availMap[n.From] = true
		availMap[n.To] = true
	}
	avail := make([]string, len(availMap))
	i := 0
	for k := range availMap {
		avail[i] = k
		i++
	}
	sort.Strings(avail)

	log.Println(nodes)
	done := make(map[string]bool)
	output := ""

	for {
		matches := make([]string, 0)
		smallestFrom := "_"
		smallestTo := "_"
		//a:
		for _, n := range nodes {
			// if !done[n.From] {
			// 	continue
			// }
			if !done[n.From] {
				for _, f := range prereq[n.From] {
					if !done[f] {
						//log.Println("skip")
						goto b
					}
				}
				if smallestFrom > n.From {
					smallestFrom = n.From
				}
			}

			if !done[n.To] {
				if smallestTo > n.To {
					smallestTo = n.To
				}
			}
		b:
			if done[n.To] {
				continue
			}

			// found it
			//log.Println("x2", n.From, n.To, done)
			for _, f := range prereq[n.To] {
				if !done[f] {
					//log.Println("skip")
					continue
				}
			}
			matches = append(matches, n.To)
		}
		//log.Println("Smallest is", smallestFrom)
		matches = []string{}
		if len(output) >= len(avail) {
			log.Fatal(output)
		}

		sort.Strings(matches)
		if true {
			//log.Println(matches)
			if len(matches) > 0 {
				//log.Fatal("adf")
			}
		}
		if len(matches) == 0 {
			//log.Println(done)

			if smallestFrom != "_" {
				matches = append(matches, smallestFrom)
			} else {
				for _, n := range nodes {
					l, r := -1, -1
					for k, v := range output {
						if string(v) == n.From {
							l = k
						}
						if string(v) == n.To {
							r = k
						}
					}
					if l >= r && l != -1 && r != -1 {
						log.Println("argh, ", n.From, n.To)
					}
				}
				output = output + smallestTo
				fmt.Println(output)
				break
				log.Fatal("no match ", output, len(output), len(nodes), smallestTo)
			}
		}
		output = output + matches[0]
		done[matches[0]] = true
		//log.Println("out", matches, output)
	}
}
