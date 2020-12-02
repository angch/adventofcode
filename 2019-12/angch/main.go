package main

import (
	"fmt"
	"log"
)

type Vector struct {
	X, Y, Z int
}
type Moon struct {
	Pos Vector
	Vel Vector
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (m Moon) Energy() int {
	p := abs(m.Pos.X) + abs(m.Pos.Y) + abs(m.Pos.Z)
	k := abs(m.Vel.X) + abs(m.Vel.Y) + abs(m.Vel.Z)
	return p * k
}

func runSim(inputs []string, steps int) {
	moons := make([]Moon, 0)
	for _, input := range inputs {
		m := Vector{}
		fmt.Sscanf(input, "<x=%d, y=%d, z=%d>", &m.X, &m.Y, &m.Z)
		// log.Println(m)
		moon := Moon{
			Pos: m,
			Vel: Vector{},
		}
		moons = append(moons, moon)
	}
	log.Println(moons)

	for s := 0; s < steps; s++ {
		for k, v := range moons {
			vel := v.Vel
			for k2, v2 := range moons {
				if k == k2 {
					continue
				}

				dx := v2.Pos.X - v.Pos.X
				if dx < -1 {
					dx = -1
				} else if dx > 1 {
					dx = 1
				}

				dy := v2.Pos.Y - v.Pos.Y
				if dy < -1 {
					dy = -1
				} else if dy > 1 {
					dy = 1
				}
				dz := v2.Pos.Z - v.Pos.Z
				if dz < -1 {
					dz = -1
				} else if dz > 1 {
					dz = 1
				}

				vel.X += dx
				vel.Y += dy
				vel.Z += dz
			}
			moons[k].Vel = vel
		}
		for k, v := range moons {
			moons[k].Pos.X += v.Vel.X
			moons[k].Pos.Y += v.Vel.Y
			moons[k].Pos.Z += v.Vel.Z
		}
	}
	log.Println(moons)

	e := 0
	for _, v := range moons {
		e += v.Energy()
	}
	log.Println(e)

}

func main() {
	inputs := []string{
		"<x=-1, y=0, z=2>",
		"<x=2, y=-10, z=-7>",
		"<x=4, y=-8, z=8>",
		"<x=3, y=5, z=-1>",
	}
	runSim(inputs, 10)

	inputs2 := []string{
		"<x=-8, y=-10, z=0>",
		"<x=5, y=5, z=10>",
		"<x=2, y=-7, z=3>",
		"<x=9, y=-8, z=-3>",
	}
	runSim(inputs2, 100)

	inputs3 := []string{
		"<x=0, y=4, z=0>",
		"<x=-10, y=-6, z=-14>",
		"<x=9, y=-16, z=-3>",
		"<x=6, y=-1, z=2>",
	}
	runSim(inputs3, 1000)
}
