package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Node struct {
	From byte
	To   byte
}

func findJob(done map[byte]bool, prereq map[byte][]byte, pending map[byte]bool) byte {
a:
	for k := byte('A'); k <= byte('Z'); k++ {
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
		return k
	}
	return 0
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

	prereq := make(map[byte][]byte)
	availMap := make(map[byte]bool)
	dot := false
	for scanner.Scan() {
		v := scanner.Text()
		a := re.FindAllStringSubmatch(v, -1)
		from, to := byte(a[0][1][0]), byte(a[0][2][0])
		//nodes = append(nodes, n)
		if prereq[from] == nil {
			prereq[from] = make([]byte, 0)
		}
		prereq[to] = append(prereq[to], from)
		if dot {
			fmt.Println(from, "->", to)
		}
		availMap[from] = true
		availMap[to] = true
	}

	//log.Println(nodes)
	done := make(map[byte]bool)
	//output := ""
	pending := make(map[byte]bool)
	for {
		job := findJob(done, prereq, pending)
		if job == 0 {
			break
		}
		done[job] = true
		fmt.Print(string(job))
	}
	fmt.Println()

	// Part 2
	done = make(map[byte]bool)
	workers := make([]int, nworker)
	workerJob := make([]byte, nworker)
	jobsdone := 0
	t := 0

	output := ""
	for t = 0; ; t++ {
		for k, w := range workers {
			if w > 0 {
				workers[k]--
			}
			if workers[k] == 0 {
				if workerJob[k] != 0 {
					pending[workerJob[k]] = false
					done[workerJob[k]] = true
					//fmt.Println("t", t, "worker", k, "done", workerJob[k])
					output = output + string(workerJob[k])
					workerJob[k] = 0
					jobsdone++
				}
			}
		}
		for k, _ := range workers {
			if workers[k] == 0 {
				job := findJob(done, prereq, pending)

				// Sanity check
				for _, v := range prereq[job] {
					if !done[v] {
						log.Fatal("Having job ", job, " before job ", v, " done")
					}
				}

				if !pending[job] && job != 0 {
					pending[job] = true
					workerJob[k] = job
				} else {
					//fmt.Println("t", t, "Worker", k, "nojob", job)
					continue
				}
				duration := int(job-'A'+1) + base
				//fmt.Println("t", t, "Worker", k, "job", job, "duration", duration, jobsdone)
				workers[k] = duration
			} else {
				//fmt.Println("t", t, "Worker", k, "doingjob", workerJob[k], workers[k])
			}
		}
		//fmt.Println("t", t, "ends jobsdone", jobsdone, output)
		if jobsdone == len(availMap) {
			break
		}
	}
	fmt.Println(t)
}
