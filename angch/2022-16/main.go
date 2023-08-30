package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Valve struct {
	FlowRate        int
	Tunnels         []string
	TunnelIndex     []int
	Destination     []int // To go to this valve, take this tunnel
	DestinationCost []int // To go to this valve, it is this far away
}

type State struct {
	ValvesOpen        uint64
	ReleasedPressure  int
	Minute            int
	PressurePerMin    int
	Location          int
	HeadingTo         int // Heading to this valve
	ElephantLocation  int
	ElephantHeadingTo int // Heading to this valve
	CameFrom          uint64
	EleCameFrom       uint64
}

var maxTime int
var allpressure int

func (s *State) Score(graph []Valve) int {
	score := (s.ReleasedPressure + s.PressurePerMin*(maxTime-s.Minute)) * 100
	score += allpressure * (maxTime - s.Minute)
	// score *= (maxTime - s.Minute + 1)

	// How close to unopened valve?
	mask := uint64(1)
	vclosed := s.ValvesOpen ^ 0xffffffffffffffff
	for _, v := range graph {
		if mask&vclosed != 0 {
			score += v.FlowRate * (maxTime - s.Minute) * 20
		}

		for _, v2 := range v.TunnelIndex {
			if (1<<v2)&vclosed != 0 {
				score += v.FlowRate * (maxTime - s.Minute) * 10
			}
		}

		mask <<= 1
	}
	return score
}

var maxValves = 80

func (s State) Clone() State {
	newState := State{
		ValvesOpen:       s.ValvesOpen,
		ReleasedPressure: s.ReleasedPressure,
		Minute:           s.Minute,
		PressurePerMin:   s.PressurePerMin,
		Location:         s.Location,
		ElephantLocation: s.ElephantLocation,
		CameFrom:         s.CameFrom,
		EleCameFrom:      s.EleCameFrom,
	}
	// copy(newState.CameFrom, s.CameFrom)
	// copy(newState.ValvesOpen, s.ValvesOpen)
	// copy(newState.EleCameFrom, s.EleCameFrom)
	return newState
}

func (s *State) Tick() {
	s.Minute++
	s.ReleasedPressure += s.PressurePerMin
}

func day16part1(graph []Valve, start, allpressure int) int {
	part1 := 0
	now := time.Now()

	states := make([]State, 0)
	states = append(states, State{
		ValvesOpen:       0,
		ReleasedPressure: 0,
		Minute:           1,
		Location:         start,
		CameFrom:         0,
		HeadingTo:        -1,
	})
	_ = states

	// Part1
	endstate := []State{}
	maxTime := 30
	bestTime, bestRate := 31, 0
	evals := 0

	for len(states) > 0 {
		// fmt.Println(len(states))
		// Pick a state.
		bestCandidate := 0
		for k, v := range states {
			if k == 0 {
				continue
			}
			if states[bestCandidate].ReleasedPressure < v.ReleasedPressure {
				bestCandidate = k
			}
			if states[bestCandidate].PressurePerMin < v.PressurePerMin {
				bestCandidate = k
			} else if states[bestCandidate].PressurePerMin == v.PressurePerMin {
				if states[bestCandidate].Minute < v.Minute {
					bestCandidate = k
				}
			}
		}
		state := states[bestCandidate]

		// states = append(states[:bestCandidate], states[bestCandidate+1:]...)
		states[bestCandidate] = states[len(states)-1]
		states = states[:len(states)-1]

		if state.Minute >= bestTime && state.ReleasedPressure+(maxTime-state.Minute)*allpressure < bestRate {
			continue
		}
		evals++
		if evals%100000 == 0 {
			fmt.Println("Eval", evals, len(states), bestTime, bestRate, state.Minute, state.ReleasedPressure)
		}

		// fmt.Println("Eval:", state.Minute, state.ReleasedPressure, state.Location, state.HeadingTo)

		// state.HeadingTo = 9
		// state.Location = 6
		// for l := state.Location; l != state.HeadingTo; {
		// 	fmt.Println("At ", l, "taking", graph[l].Destination[state.HeadingTo])
		// 	l = graph[l].Destination[state.HeadingTo]
		// }
		// if true {
		// 	log.Fatal("s")
		// }

		if state.Minute == maxTime {
			// endstate = append(endstate, state)
			if state.ReleasedPressure > bestRate {
				bestRate = max(bestRate, state.ReleasedPressure)
				bestTime = state.Minute
				endstate = append(endstate, state)
			} else if state.ReleasedPressure == bestRate {
				bestTime = min(bestTime, state.Minute)
				endstate = append(endstate, state)
			}
			// fmt.Println("Late end", state.Minute, state.ReleasedPressure, len(states), bestRate, bestTime)
			continue
		}

		// log.Printf("State: %+v\n", state)
		if state.PressurePerMin == allpressure {
			mleft := maxTime - state.Minute
			state.ReleasedPressure += mleft * allpressure
			// fmt.Println("Early end", state.Minute, state.ReleasedPressure, len(states), bestRate, bestTime)

			if state.ReleasedPressure > bestRate {
				bestRate = max(bestRate, state.ReleasedPressure)
				bestTime = state.Minute
				state.Minute = maxTime
				endstate = append(endstate, state)
			} else if state.ReleasedPressure == bestRate {
				bestTime = min(bestTime, state.Minute)
				state.Minute = maxTime
				endstate = append(endstate, state)
			}

			// state.Minute = maxTime
			// endstate = append(endstate, state)
			continue
		}

		// Is valve open?
		l := state.Location

		vo := (1 << l) & state.ValvesOpen
		if vo == 0 && graph[l].FlowRate > 0 {
			// fmt.Println("Opening valve")
			// newState := state.Clone()
			newState := state
			newState.ValvesOpen |= (1 << l)
			newState.PressurePerMin += graph[l].FlowRate
			newState.CameFrom = 1 << l
			newState.Tick()
			states = append(states, newState)
			// continue
		} else if vo == 1 && state.HeadingTo == l {
			log.Fatal("what")
		}

		// Where can we go?
		// for _, v := range graph[l].TunnelIndex {
		// 	cf := (1 << v) & state.CameFrom
		// 	if cf != 0 {
		// 		continue
		// 	}
		// 	// fmt.Println("Moving to", v, state.Minute, state.ReleasedPressure)
		// 	newState := state
		// 	newState.Location = v
		// 	newState.CameFrom |= (1 << l)
		// 	newState.Tick()
		// 	states = append(states, newState)
		// }

		// Where to?
		headTo := []int{}
		if state.HeadingTo == -1 || state.HeadingTo == l {
			closedValves := ^state.ValvesOpen
			for destId, v := range graph {
				if v.FlowRate == 0 {
					continue
				}
				if destId == l {
					continue
				}
				if (1<<destId)&closedValves != 0 {
					if (1<<destId)&state.CameFrom == 0 {
						headTo = append(headTo, destId)
					}
				}
			}
			// fmt.Println("Heading to", headTo, "from valve", l)
		} else {
			headTo = []int{state.HeadingTo}
		}

		// fmt.Println("Heading to", headTo, "from valve", l)
		for _, v := range headTo {
			tunnelId := graph[l].Destination[v]
			if tunnelId == -1 {
				continue
			}
			newState := state
			newState.Location = tunnelId
			newState.HeadingTo = v
			newState.CameFrom |= (1 << l)
			newState.Tick()
			// log.Println("At valve", l, "heading to", v, "using tunnel", tunnelId, "minute", newState.Minute, "pressure", newState.ReleasedPressure, newState.ValvesOpen)
			states = append(states, newState)
			// break
		}

	}
	// fmt.Println("endstate", len(endstate), allpressure)
	maxFlow := 0
	for _, v := range endstate {
		maxFlow = max(maxFlow, v.ReleasedPressure)
	}
	// fmt.Println("Maxflow", maxFlow)
	// part1
	part1 = maxFlow
	fmt.Println("Part1: Time taken", time.Since(now), evals, "evals taken, best time", bestTime)

	return part1
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    State // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func day16part2(graph []Valve, start, allpressure int) int {
	part2 := 0
	now := time.Now()

	// states := make([]State, 0)
	initState := State{
		ValvesOpen:       0,
		ReleasedPressure: 0,
		Minute:           1,
		Location:         start,
		ElephantLocation: start,
		CameFrom:         0,
		EleCameFrom:      0,
	}

	endstate := []State{}
	maxTime = 26
	bestTime, bestRate := 31, 0

	evals := 0
	humanMoves := make([]State, 0, 10)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	pq.Push(&Item{value: initState, priority: 0})
	skipped := 0

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		state := item.value

		if state.Minute > bestTime && state.ReleasedPressure+(maxTime-state.Minute)*allpressure < bestRate {
			skipped++
			continue
		}
		evals++
		if evals%100000 == 0 {
			fmt.Println("Eval", evals, len(pq), bestRate, bestTime, "cur:", state.ReleasedPressure, state.Minute, item.priority, skipped)
		}

		if state.Minute == maxTime {
			// endstate = append(endstate, state)
			if state.ReleasedPressure > bestRate {
				bestRate = max(bestRate, state.ReleasedPressure)
				bestTime = state.Minute
				endstate = append(endstate, state)
			} else if state.ReleasedPressure == bestRate {
				bestTime = min(bestTime, state.Minute)
				endstate = append(endstate, state)
			}
			// fmt.Println("Late end", state.Minute, state.ReleasedPressure, len(states), bestRate, bestTime)
			continue
		}

		if state.PressurePerMin == allpressure {
			mleft := maxTime - state.Minute
			state.ReleasedPressure += mleft * allpressure
			// fmt.Println("Early end", state.Minute, state.ReleasedPressure, len(states), bestRate, bestTime)

			if state.ReleasedPressure > bestRate {
				bestRate = max(bestRate, state.ReleasedPressure)
				bestTime = state.Minute
				state.Minute = maxTime
				endstate = append(endstate, state)
			} else if state.ReleasedPressure == bestRate {
				bestTime = min(bestTime, state.Minute)
				state.Minute = maxTime
				endstate = append(endstate, state)
			}
			continue
		}

		humanMoves := humanMoves[:0]

		// Is valve open?
		l := state.Location
		vo := (1 << l) & state.ValvesOpen
		if vo == 0 && graph[l].FlowRate > 0 {
			// fmt.Println("Opening valve")
			newState := state.Clone()

			newState.ValvesOpen |= 1 << l
			// newState.CameFrom = make([]bool, maxValves)
			// newState.EleCameFrom = make([]bool, maxValves)
			newState.CameFrom = 1 << l
			// newState.EleCameFrom = 0
			newState.PressurePerMin += graph[l].FlowRate
			humanMoves = append(humanMoves, newState)
		}

		// Where can we go?
		for _, v := range graph[l].TunnelIndex {
			cf := (1 << v) & state.CameFrom
			if cf != 0 {
				continue
			}
			// fmt.Println("Moving to", v, state.Minute, state.ReleasedPressure)
			newState := state.Clone()

			newState.Location = v
			newState.CameFrom |= 1 << l
			humanMoves = append(humanMoves, newState)
		}

		// Elephant moves
		for _, state2 := range humanMoves {
			l := state2.ElephantLocation

			vo := (1 << l) & state2.ValvesOpen
			if vo == 0 && graph[l].FlowRate > 0 {
				// fmt.Println("Opening valve")
				newState := state.Clone()
				newState.ValvesOpen |= 1 << l
				newState.EleCameFrom = 1 << l
				newState.PressurePerMin += graph[l].FlowRate
				newState.Tick()
				pq.Push(&Item{value: newState, priority: newState.Score(graph)})
				// states = append(states, newState)
			}

			// Where can we go?
			for _, v := range graph[l].TunnelIndex {
				cf := (1 << v) & state2.EleCameFrom
				if cf != 0 {
					continue
				}
				// fmt.Println("Moving to", v, state.Minute, state.ReleasedPressure)
				newState := state.Clone()

				newState.ElephantLocation = v
				newState.EleCameFrom |= 1 << l
				newState.Tick()
				// states = append(states, newState)
				pq.Push(&Item{value: newState, priority: newState.Score(graph)})
			}
		}

	}
	// fmt.Println("endstate", len(endstate), allpressure)
	maxFlow := 0
	for _, v := range endstate {
		maxFlow = max(maxFlow, v.ReleasedPressure)
	}
	// fmt.Println("Maxflow", maxFlow)
	// part1
	part2 = maxFlow
	fmt.Println("Time taken", time.Since(now))

	return part2
}

func RouteTo(from, to int, graph []Valve, visited uint64) int {
	if from == to {
		return 0
	}
	for _, v := range graph[from].TunnelIndex {
		if v == to {
			return 1
		}
	}
	best := 9999999999
	for _, v := range graph[from].TunnelIndex {
		if visited&(1<<v) == 0 {
			best = min(best, 1+RouteTo(v, to, graph, visited|(1<<v)))
		}
	}
	return best
}

func day16(file string) (int, int) {
	part1, part2 := 0, 0
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	valveIndex := make(map[string]int)
	graph := make([]Valve, 0)

	// allpressure := 0
	start := 0
	for scanner.Scan() {
		t := scanner.Text()
		v := Valve{}
		name := ""
		l, r, ok := strings.Cut(t, "to valves ")
		if !ok {
			l, r, ok = strings.Cut(t, "to valve ")
			// log.Fatal("s", t, l, r, ok)
			if !ok {
				continue
			}
		}
		fmt.Sscanf(l, "Valve %s has flow rate=%d;", &name, &v.FlowRate)
		allpressure += v.FlowRate

		tun := strings.Split(r, ", ")
		v.Tunnels = append(v.Tunnels, tun...)

		valveIndex[name] = len(graph)
		if name == "AA" {
			start = len(graph)
		}
		graph = append(graph, v)
	}
	for k := range graph {
		for _, v2 := range graph[k].Tunnels {
			graph[k].TunnelIndex = append(graph[k].TunnelIndex, valveIndex[v2])
		}
	}

	// routeTo :=

	// Map out the destinations
	// For every node's as a destination, we mark which tunnel to take, -1 if already there
	for tunnelId, v := range graph {
		// The choices are the ones in TunnelIndex, or -1
		destination := make([]int, len(graph))
		cost := make([]int, len(graph))

		for destTunnelId := range destination {
			if destTunnelId == tunnelId {
				// We're already here
				destination[destTunnelId] = -1
				cost[destTunnelId] = 0
				continue
			}

			// Find the shortest path
			bestTunnel := -1
			bestScore := 99999
			for _, v2 := range v.TunnelIndex {
				score := 1 + RouteTo(v2, destTunnelId, graph, 1<<tunnelId)
				if tunnelId == 3 && destTunnelId == 9 {
					log.Println("Going", v2, destTunnelId, score)
				}
				if score <= bestScore {
					bestScore = score
					bestTunnel = v2
				}
			}
			destination[destTunnelId] = bestTunnel
			cost[destTunnelId] = bestScore
			// fmt.Println("Finding shortest path")
		}

		graph[tunnelId].Destination = destination
		graph[tunnelId].DestinationCost = cost
	}
	fmt.Printf("Graph: %+v\n", graph)
	for k, v := range graph {
		fmt.Printf(" %d %s %d %d %d\n", k, v.Tunnels, v.TunnelIndex, v.Destination, v.DestinationCost)
	}
	// if true {
	// 	return 1, 2
	// }

	part1 = day16part1(graph, start, allpressure)
	// part1 = 1651
	fmt.Println("part1", part1)
	// part2 = day16part2(graph, start, allpressure)

	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, part2 := day16("test.txt")

	// part1 Time taken 13.287075ms 139369 evals taken
	// Time taken 1.089914894s 6049832 evals taken <- dest
	fmt.Println(part1, part2)
	if part1 != 1651 || part2 != 1707 {
		log.Fatal("Bad test expect 1651 and 1707")
	}
	// fmt.Println(day16("test2.txt")) // 1563 too low

	// part1 Time taken 1m26.050488254s 404526530 evals taken
	// fmt.Println(day16("input.txt")) // 1563 too low
}
