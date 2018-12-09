package main

import (
	"container/list"
	"fmt"
	"log"
)

// 3165687872

func main() {
	nplayers, target := 452, 70784

	l := list.New()

	current := l.PushFront(0)
	currentMarble := 1
	currentPlayer := 0

	scores := make([]int, nplayers+1)
	for {

		currentPlayer++
		if currentPlayer > nplayers {
			currentPlayer = 1
		}

		if currentMarble%23 == 0 {
			score := currentMarble

			for i := 0; i < 8; i++ {
				current = current.Prev()
				if current == nil {
					current = l.Back()
				}
			}
			remove := current
			current = current.Next()
			if current == nil {
				current = l.Front()
			}
			v := remove.Value
			switch v.(type) {
			case int:
				score += v.(int)
			}

			//score += int(remove.Value)
			//log.Println(score)
			l.Remove(remove)
			scores[currentPlayer] += score
		} else {
			current = l.InsertAfter(currentMarble, current)
		}
		current = current.Next()
		if current == nil {
			current = l.Front()
		}

		if false {
			fmt.Print("[", currentPlayer, "]")
			for v := l.Front(); v != nil; v = v.Next() {
				fmt.Print(" ", v.Value)
			}
			fmt.Println("")
		}

		currentMarble++

		if currentMarble == target {
			break
		}
	}
	bestScore, bestPlayer := 0, 0
	for k, v := range scores {
		if v > bestScore {
			bestScore = v
			bestPlayer = k
		}
	}
	log.Println(bestPlayer, bestScore)
}
