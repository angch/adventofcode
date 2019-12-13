package main

import (
	"log"
	"sync"
)

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func AmpRunner(program map[int]int, part int) int {
	var signals [][]int
	if part == 1 {
		signals = permutations([]int{0, 1, 2, 3, 4})
	} else {
		signals = permutations([]int{5, 6, 7, 8, 9})
	}

	result := make(chan int)
	for _, signal := range signals {
		go func(signal []int) {
			signalchan := make([](chan int), 6)
			for i := range signalchan {
				signalchan[i] = make(chan int, 1)
			}

			wg := sync.WaitGroup{}
			wg.Add(5)

			// Wire up all the "Amp"s and connect the inputs together.
			// Note all of them will start running until blocked and waiting
			// on their inputs (phase first):
			go func() { IntCodeVM(program, signalchan[0], signalchan[1]); wg.Done() }()
			go func() { IntCodeVM(program, signalchan[1], signalchan[2]); wg.Done() }()
			go func() { IntCodeVM(program, signalchan[2], signalchan[3]); wg.Done() }()
			go func() { IntCodeVM(program, signalchan[3], signalchan[4]); wg.Done() }()
			go func() { result <- IntCodeVM(program, signalchan[4], signalchan[0]); wg.Done() }()

			// Load all phases into the inputs to the Amps
			for i := range signal {
				signalchan[i] <- signal[i]
			}
			// Signal everyone to start processing
			signalchan[0] <- 0
			wg.Wait()
		}(signal)
	}
	max := 0
	for range signals {
		r := <-result
		if r > max {
			max = r
		}
	}
	log.Println("max", max)
	return max
}
