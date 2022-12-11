package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type person struct {
	HP     int
	Mana   int
	Damage int
	Armor  int
}

type effectDuration struct {
	shield   int
	poison   int
	recharge int
}

type gamestate struct {
	turns    int
	player   person
	boss     person
	effect   effectDuration
	win      bool
	lose     bool
	manaused int
	hard     bool
}

var actioncost = map[string]int{
	"Magic Missile": 53,
	"Drain":         73,
	"Shield":        113,
	"Poison":        173,
	"Recharge":      229,
}

func doTurn(state gamestate, action string, verbose bool) gamestate {
	if state.turns == 1 {
		if verbose {
			fmt.Println("-- Player turn --")
			fmt.Printf("- Player has %d hit points, %d armor, %d mana\n", state.player.HP, state.player.Armor, state.player.Mana)
			fmt.Printf("- Boss has %d hit points\n", state.boss.HP)
		}
		if state.hard {
			state.player.HP--
			if state.player.HP <= 0 {
				state.lose = true
				return state
			}
		}
	}

	state.player.Mana -= actioncost[action]
	state.manaused += actioncost[action]
	switch action {
	case "Magic Missile":
		state.boss.HP -= 4
		if verbose {
			fmt.Printf("Player casts %s, dealing 4 damage.\n", action)
		}
	case "Drain":
		state.boss.HP -= 2
		state.player.HP += 2
		if verbose {
			fmt.Printf("Player casts %s, dealing 2 damage, and healing 2 hit points.\n", action)
		}
	case "Shield":
		state.effect.shield = 6
		if verbose {
			fmt.Println("Player casts Shield, increasing armor by 7.")
		}
	case "Poison":
		state.effect.poison = 6
		if verbose {
			fmt.Printf("Player casts %s.\n", action)
		}
	case "Recharge":
		state.effect.recharge = 5
		if verbose {
			fmt.Printf("Player casts %s.\n", action)
		}
	default:
		panic("No such action " + action)
	}

	state.turns++
	if state.boss.HP <= 0 {
		// Immediate win from state effects (Magic The Gathering term)
		state.win = true
		if verbose {
			fmt.Println("This kills the boss, and the player wins.")
		}
		return state
	}

	// boss turn
	if verbose {
		fmt.Println("\n-- Boss turn --")
		fmt.Printf("- Player has %d hit points, %d armor, %d mana\n", state.player.HP, state.player.Armor, state.player.Mana)
		fmt.Printf("- Boss has %d hit points\n", state.boss.HP)
	}
	// if state.hard {
	// 	state.player.HP--
	// 	if state.player.HP <= 0 {
	// 		state.lose = true
	// 		return state
	// 	}
	// }
	state.player.Armor = 0
	if state.effect.shield > 0 {
		state.effect.shield--
		state.player.Armor = 7
		if verbose {
			fmt.Printf("Shield's timer is now %d.\n", state.effect.shield)
		}
	}
	if state.effect.poison > 0 {
		state.effect.poison--
		state.boss.HP -= 3
		if verbose {
			fmt.Printf("Poison deals %d damage; its timer is now %d.\n", 3, state.effect.poison)
		}
	}
	if state.effect.recharge > 0 {
		state.effect.recharge--
		state.player.Mana += 101
		if verbose {
			fmt.Printf("Recharge provides %d mana; its timer is now %d.\n", 101, state.effect.recharge)
		}
	}
	state.turns++
	if state.boss.HP <= 0 {
		// Immediate win from state effects (Magic The Gathering term)
		state.win = true
		if verbose {
			fmt.Println("This kills the boss, and the player wins.")
		}
		return state
	}

	damage := state.boss.Damage - state.player.Armor
	if damage < 1 {
		damage = 1
	}
	state.player.HP -= damage
	if verbose {
		fmt.Printf("Boss attacks for %d damage!\n", damage)
	}
	if state.player.HP <= 0 {
		// Immediate lose from state effects (Magic The Gathering term)
		state.lose = true
		return state
	}
	if verbose {
		fmt.Println()
	}

	// player turn
	if verbose {
		fmt.Println("-- Player turn --")
		fmt.Printf("- Player has %d hit points, %d armor, %d mana\n", state.player.HP, state.player.Armor, state.player.Mana)
		fmt.Printf("- Boss has %d hit points\n", state.boss.HP)
	}
	if state.hard {
		state.player.HP--
		if state.player.HP <= 0 {
			state.lose = true
			return state
		}
	}
	state.player.Armor = 0
	if state.effect.shield > 0 {
		state.effect.shield--
		state.player.Armor = 7
		if verbose {
			fmt.Printf("Shield's timer is now %d.\n", state.effect.shield)
		}
	}
	if state.effect.poison > 0 {
		state.effect.poison--
		state.boss.HP -= 3
		if verbose {
			fmt.Printf("Poison deals %d damage; its timer is now %d.\n", 3, state.effect.poison)
		}
	}
	if state.effect.recharge > 0 {
		state.effect.recharge--
		state.player.Mana += 101
		if verbose {
			fmt.Printf("Recharge provides %d mana; its timer is now %d.\n", 101, state.effect.recharge)
		}
	}
	state.turns++
	if state.boss.HP <= 0 {
		// Immediate win from state effects (Magic The Gathering term)
		state.win = true
		return state
	}

	return state
}

func day21(file string, player person, actions []string) (int, int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	boss := person{}

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	_, r, _ := strings.Cut(scanner.Text(), ": ")
	hp, err := strconv.Atoi(r)
	if err != nil {
		log.Fatal(err)
	}
	scanner.Scan()
	_, r, _ = strings.Cut(scanner.Text(), ": ")
	damage, _ := strconv.Atoi(r)
	scanner.Scan()

	boss.HP = hp
	boss.Damage = damage

	best, besthard := 99999999, 99999999
	state := gamestate{boss: boss, player: player, turns: 1}
	if len(actions) > 0 {
		for _, v := range actions {
			state = doTurn(state, v, true)
		}
	} else {
		oldstate := state
		queue := []gamestate{state}
		now := time.Now()

		for len(queue) > 0 {
			state := queue[0]
			queue = queue[1:]

			if time.Since(now) > 1*time.Second {
				now = time.Now()
				fmt.Println("Queue len is", len(queue), "and state is", state.turns, "turns in, best is", besthard)
			}
			if state.turns > 50 {
				continue
			}

			actions := []string{}
			for k, v := range actioncost {
				if state.player.Mana >= v && v+state.manaused < best {
					actions = append(actions, k)
				}
			}

			for _, v := range actions {
				newstate := doTurn(state, v, false)
				if newstate.win {
					if newstate.manaused < best {
						best = newstate.manaused
					}
				} else if newstate.lose {

				} else {
					queue = append(queue, newstate)
				}
			}
		}

		state = oldstate
		state.hard = true
		queue = []gamestate{state}
		// state.player.HP--
		now = time.Now()
		// fmt.Printf("State %+v", state)
		for len(queue) > 0 {
			state = queue[0]
			queue = queue[1:]

			if time.Since(now) > time.Second {
				now = time.Now()
				fmt.Println("Queue len is", len(queue), "and state is", state.turns, "turns in, best is", besthard)
			}

			actions := []string{}
			for k, v := range actioncost {
				if state.player.Mana >= v && v+state.manaused <= besthard {
					actions = append(actions, k)
				}
			}

			for _, v := range actions {
				newstate := doTurn(state, v, false)
				if newstate.win {
					if newstate.manaused < besthard {
						besthard = newstate.manaused
					}
				} else if newstate.lose {

				} else {
					queue = append(queue, newstate)
				}
			}
		}

	}
	_ = state

	return best, besthard
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// fmt.Println(day21("test.txt", person{HP: 10, Mana: 250}, []string{"Poison", "Magic Missile"}))
	// fmt.Println(day21("test2.txt", person{HP: 10, Mana: 250}, []string{"Recharge", "Shield", "Drain", "Poison", "Magic Missile"}))
	fmt.Println(day21("input.txt", person{HP: 50, Mana: 500}, []string{}))
	// 1309 too high
}
