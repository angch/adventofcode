package main

import "log"

// Damn Go and slices
func insertAt(a []int, insertAfterPos, val int) []int {
	//log.Println("before:", a, insertAfterPos, val)
	// there's probably an easier way, meh.
	//ar = append(ar, 0)
	a = append(a[:insertAfterPos], append([]int{val}, a[insertAfterPos:]...)...)
	//log.Println("after:", a)
	return a
}

func main() {
	nplayers := 452
	target := 7078400
	//boom := 8317 // sanity while testing
	//boom := 384892

	circles := make([]int, 0)
	circles = append(circles, 0)
	currentMarble := 1
	currentPlayer := 0
	lastMarblePos := 1
	turn := 1

	log.Println("[", currentPlayer, "]", circles, currentMarble, lastMarblePos)
	scores := make([]int, nplayers+1)

	for {
		//insertAt()
		currentPlayer++
		if currentPlayer > nplayers {
			currentPlayer = 1
		}

		score := 0
		if currentMarble%23 == 0 {
			//continue
			score = currentMarble
			take := lastMarblePos - 9
			for take < 0 {
				take += len(circles)
			}
			//log.Println("take", take, circles[take])
			score += circles[take]

			copy(circles[take:], circles[take+1:])
			circles = circles[:len(circles)-1]
			lastMarblePos = take
			// if lastMarblePos < 0 {
			// 	lastMarblePos = len(circles)
			// }
			//currentMarble++
			//currentPlayer--
			//log.Println("score is", score)
			//scores = append(scores, score)
			// if score == boom {
			// 	// Magic score multiplier
			// 	score -= currentMarble
			// 	score += currentMarble * 100
			// 	break
			// }

			scores[currentPlayer] += score

			// if score > boom {
			// 	break
			// }
			if score == target {
				break
			}
			//continue
		} else {
			circles = insertAt(circles, lastMarblePos, currentMarble)
		}
		lastMarblePos = (lastMarblePos+1)%(len(circles)) + 1
		//log.Println("[", currentPlayer, "]", circles, currentMarble, lastMarblePos)
		currentMarble++
		if currentMarble > target {
			break
		}
		turn++

	}
	bestScore := 0
	winningElf := 0
	for i, s := range scores {
		if s > bestScore {
			bestScore = s
			winningElf = i
		}
	}
	log.Println(scores, bestScore, winningElf)
	log.Println(scores, scores[236], winningElf)
	if false {
		log.Println(turn, lastMarblePos, currentPlayer, currentMarble, nplayers, target)
	}
}
