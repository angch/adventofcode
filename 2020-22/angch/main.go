package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Hands map[int][]int

func (h *Hands) Hash() string {
	j := md5.New()
	for _, v := range (*h)[1] {
		j.Write([]byte{byte(v)})
	}
	for _, v := range (*h)[2] {
		j.Write([]byte{byte(v) + 64})
	}
	k := j.Sum([]byte{})
	return string(k[:])
}

func do(fileName string) (ret1 int, ret2 int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	hands := make(map[int][]int, 0)

	for scanner.Scan() {
		l := scanner.Text()
		_ = l

		player := 0
		fmt.Sscanf(l, "Player %d:", &player)
		hands[player] = make([]int, 0)

		for scanner.Scan() {
			l := scanner.Text()
			_ = l

			if l == "" {
				break
			}
			i, _ := strconv.Atoi(l)

			hands[player] = append(hands[player], i)
		}
	}

a:
	for {
		top := make([]int, len(hands)+1)
		for k, v := range hands {
			if len(v) == 0 {
				break a
			}
			top[k] = v[0]
			hands[k] = v[1:]
		}

		if top[1] > top[2] {
			hands[1] = append(hands[1], top[1], top[2])
		} else {
			hands[2] = append(hands[2], top[2], top[1])
		}
		// log.Println(hands)
	}

	for k, v := range hands {
		if len(v) == 0 {
			continue
		}
		log.Println("Winner", k, hands)

		for k2, v2 := range v {
			ret1 += v2 * (len(v) - k2)
		}
	}
	return ret1, ret2
}

var debug = false
var history2 = map[string]int{}

func subgame(hands Hands, round int, game int) (int, Hands) {
	winner := 0
	subgamecount := game

	sum2 := hands.Hash()
	if h3, ok := history2[sum2]; ok {
		if debug {
			log.Println("cached", len(history2), round, game)
		}
		return h3, hands
	}
	history := make(map[string]bool)
a:
	for {
		if debug {
			log.Println("Round", round, "Game", game)
			log.Println(hands)
		}
		round++
		sum := hands.Hash()
		if history[sum] {
			if debug {
				log.Println("inf game", round, game)
			}
			winner = 1
			break
		}

		top := make([]int, len(hands)+1)
		history[sum] = true
		for k, v := range hands {
			if len(v) == 0 {
				if k == 1 {
					winner = 2
				} else {
					winner = 1
				}
				break a
			}
			top[k] = v[0]
			hands[k] = v[1:]
		}

		if debug {
			log.Println("top", top[1], len(hands[1]), top[2], len(hands[2]))
		}

		if top[1] <= len(hands[1]) && top[2] <= len(hands[2]) {
			h2 := make(Hands)
			h2[1] = make([]int, top[1])
			h2[2] = make([]int, top[2])
			copy(h2[1], hands[1])
			copy(h2[2], hands[2])
			if debug {
				log.Println("Playing subgame", h2)
			}
			subgamecount++
			w2, _ := subgame(h2, 1, subgamecount)
			if debug {
				log.Println("Winner of subgame is", w2)
			}

			if w2 == 1 {
				hands[1] = append(hands[1], top[1], top[2])
			} else {
				hands[2] = append(hands[2], top[2], top[1])
			}
			continue
		}

		if top[1] > top[2] {
			hands[1] = append(hands[1], top[1], top[2])
		} else {
			hands[2] = append(hands[2], top[2], top[1])
		}
	}
	history2[sum2] = winner

	return winner, hands
}

func do2(fileName string) (ret1 int, ret2 int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	hands := make(Hands)

	for scanner.Scan() {
		l := scanner.Text()
		_ = l

		player := 0
		fmt.Sscanf(l, "Player %d:", &player)
		hands[player] = make([]int, 0)

		for scanner.Scan() {
			l := scanner.Text()
			_ = l

			if l == "" {
				break
			}
			i, _ := strconv.Atoi(l)

			hands[player] = append(hands[player], i)
		}
	}

	w, hands := subgame(hands, 1, 1)
	log.Println("Winner", w, hands)
	for _, v := range hands {
		if len(v) == 0 {
			continue
		}

		for k2, v2 := range v {
			ret1 += v2 * (len(v) - k2)
			// log.Println(v2, (len(v) - k2))
		}
	}
	// ret1 = count
	// log.Println(history2)
	ret2 = len(history2)

	return ret1, ret2
}

func main() {
	// log.Println(do("test.txt"))
	// log.Println(do2("test.txt"))
	log.Println(do("input.txt"))
	log.Println(do2("input.txt"))
}
