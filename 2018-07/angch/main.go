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

// 653
// 894
// 893

type Node struct {
	From string
	To   string
}

func findJob(nodes []Node, done map[string]bool, prereq map[string][]string, pending map[string]bool) string {
	//	smallest := "_"
a:
	for kk := 'A'; kk <= 'Z'; kk++ {
		k := string(kk)
		if done[k] {
			continue
		}
		if pending[k] {
			continue
		}

		for _, f := range prereq[k] {
			if !done[f] {
				continue a
			}
			if pending[f] {
				continue a
			}
		}
		return string(k)
	}
	return "_"
}

func findJobOld(nodes []Node, done map[string]bool, prereq map[string][]string, pending map[string]bool) string {
	smallestFrom := "_"
	smallestTo := "_"
a:
	for _, n := range nodes {
		if !done[n.From] {
			if pending[n.From] {
				continue
			}
			for _, f := range prereq[n.From] {
				if !done[f] {
					continue a
				}
				if pending[f] {
					continue a
				}
			}
			if smallestFrom > n.From {
				smallestFrom = n.From
			}
		}
		if !done[n.To] {
			if pending[n.To] {
				continue
			}
			for _, f := range prereq[n.To] {
				if !done[f] {
					continue a
				}
				if pending[f] {
					continue a
				}
			}
			if smallestTo > n.To {
				smallestTo = n.To
			}
		}
	}
	//fmt.Println("smallest ", smallestFrom, smallestTo)

	if smallestFrom != "_" {
		return smallestFrom
	}
	return smallestTo
}

func main() {
	nworker := 2
	base := 0
	fileName := "input2.txt"
	nworker, base, fileName = 5, 60, "input.txt"

	file, err := os.Open(fileName)
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
	//output := ""
	pending := make(map[string]bool)
	for {
		job := findJob(nodes, done, prereq, pending)
		if job == "_" {
			break
		}
		done[job] = true
		fmt.Print(job)
	}
	fmt.Println()

	// Part 2
	done = make(map[string]bool)
	workers := make([]int, nworker)
	workerJob := make([]string, nworker)
	jobsdone := 0
	t := 0

	output := ""
	for t = 0; jobsdone < len(avail); t++ {
		for k, w := range workers {
			if w > 0 {
				workers[k]--
			}
			if workers[k] == 0 {
				if workerJob[k] != "" {
					pending[workerJob[k]] = false
					done[workerJob[k]] = true
					fmt.Println("t", t, "worker", k, "done", workerJob[k])
					output = output + workerJob[k]
					workerJob[k] = ""

					jobsdone++
				}
			}
		}
		for k, _ := range workers {
			if workers[k] == 0 {
				job := findJob(nodes, done, prereq, pending)

				// Sanity check
				for _, v := range prereq[job] {
					if !done[v] {
						log.Fatal("Having job ", job, " before job ", v, " done")
					}
				}

				if !pending[job] && job != "_" {
					//done[job] = true
					pending[job] = true
					workerJob[k] = job
				} else {

					fmt.Println("t", t, "Worker", k, "nojob", job)
					continue
				}
				duration := int(byte(job[0])-'A'+1) + base
				fmt.Println("t", t, "Worker", k, "job", job, "duration", duration, jobsdone)
				workers[k] = duration
			} else {
				fmt.Println("t", t, "Worker", k, "doingjob", workerJob[k], workers[k])
			}
		}
		fmt.Println("t", t, "ends jobsdone", jobsdone, output)
		if jobsdone == len(avail) {
			fmt.Println("bre")
			break
		}
	}
	//t--
	//log.Println(t, workers)
	//t += max
	fmt.Println(t)
}
