package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Valve struct {
	FlowRate    int
	Tunnels     []string
	TunnelIndex []int
}

type State struct {
	ValvesOpen       []bool
	ReleasedPressure int
	Minute           int
	PressurePerMin   int
	Location         int
	ElephnatLocation int
	CameFrom         []bool
}

var maxValves = 80

func (s State) Clone() State {
	newState := State{
		ValvesOpen:       make([]bool, len(s.ValvesOpen), maxValves),
		ReleasedPressure: s.ReleasedPressure,
		Minute:           s.Minute,
		PressurePerMin:   s.PressurePerMin,
		Location:         s.Location,
		CameFrom:         make([]bool, len(s.CameFrom), maxValves),
	}
	// for k, v := range s.ValvesOpen {
	// 	newState.ValvesOpen[k] = v
	// }
	// for k, v := range s.CameFrom {
	// 	newState.CameFrom[k] = v
	// }
	copy(newState.CameFrom, s.CameFrom)
	copy(newState.ValvesOpen, s.ValvesOpen)
	return newState
}

func (s *State) Tick() {
	s.Minute++
	s.ReleasedPressure += s.PressurePerMin
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

	allpressure := 0
	start := 0
	now := time.Now()
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
	fmt.Printf("Graph: %+v\n", graph)

	states := make([]State, 0)
	states = append(states, State{
		ValvesOpen:       make([]bool, maxValves),
		ReleasedPressure: 0,
		Minute:           1,
		Location:         start,
		CameFrom:         make([]bool, maxValves),
	})
	_ = states

	// Part1
	endstate := []State{}
	maxTime := 30
	bestTime, bestRate := 31, 0

	for len(states) > 0 {
		// fmt.Println(len(states))
		// Pick a state.
		bestCandidate := 0
		for k, v := range states {
			if k == 0 {
				continue
			}
			if states[bestCandidate].PressurePerMin < v.PressurePerMin {
				bestCandidate = k
			} else if states[bestCandidate].PressurePerMin == v.PressurePerMin {
				if states[bestCandidate].Minute < v.Minute {
					bestCandidate = k
				}
			}
			// if len(states[bestCandidate].ValvesOpen) < len(v.ValvesOpen) {
			// 	bestCandidate = k
			// } else if len(states[bestCandidate].ValvesOpen) == len(v.ValvesOpen) {
			// 	// if states[bestCandidate].Minute < v.Minute {
			// 	// 	bestCandidate = k
			// 	// } else if states[bestCandidate].Minute == v.Minute {
			// 	if states[bestCandidate].PressurePerMin < v.PressurePerMin {
			// 		bestCandidate = k
			// 	}
			// 	// }
			// }

		}
		state := states[bestCandidate]

		// states = append(states[:bestCandidate], states[bestCandidate+1:]...)
		states[bestCandidate] = states[len(states)-1]
		states = states[:len(states)-1]

		if state.Minute > bestTime && state.ReleasedPressure+(maxTime-state.Minute)*allpressure < bestRate {
			continue
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
		if !state.ValvesOpen[l] && graph[l].FlowRate > 0 {
			// fmt.Println("Opening valve")
			newState := state.Clone()
			newState.ValvesOpen[l] = true
			newState.CameFrom = make([]bool, maxValves)
			newState.PressurePerMin += graph[l].FlowRate
			newState.Tick()
			states = append(states, newState)
		}

		// Where can we go?
		for _, v := range graph[l].TunnelIndex {
			if state.CameFrom[v] {
				continue
			}
			// fmt.Println("Moving to", v, state.Minute, state.ReleasedPressure)
			newState := state.Clone()
			newState.Location = v
			newState.CameFrom[l] = true
			newState.Tick()
			states = append(states, newState)
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
	fmt.Println("Time taken", time.Since(now))

	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, part2 := day16("test.txt")

	fmt.Println(part1, part2)
	if part1 != 1651 || part2 != 0 {
		log.Fatal("Bad test")
	}
	// fmt.Println(day16("test2.txt")) // 1563 too low
	fmt.Println(day16("input.txt")) // 1563 too low
}
