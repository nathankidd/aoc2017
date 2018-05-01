package main

import "fmt"
import "os"
import "bufio"
import "sort"

type Point struct {
	x int
	y int
	z int
}

type Particle struct {
	pos Point
	vel Point
	acc Point
}

type Pair struct {
	order int
	val   int
}

func loose_order_match(n int, a []Pair, b []Pair) bool {
	if a[n].order == b[n].order {
		return true
	}
	// Having the same value in an adjacent position is close enough
	// Check lower array positions
	for i := n; i >= 0 && b[n].val == b[i].val; i-- {
		if a[n].order == b[i].order {
			return true
		}
	}
	// And later array positions
	for i := n; i < len(b) && b[n].val == b[i].val; i++ {
		if a[n].order == b[i].order {
			return true
		}
	}
	return false
}

// Check for order
// Once the particle with greatest or equal acceleration also
// has greatest or equal velocity and greatest or equal position
// for a given plane it is no longer possible to have a collision.
// X plane picked here, any can work
func particles_sorted(field map[int]Particle) bool {
	var pos []Pair
	var vel []Pair
	var acc []Pair
	for num, p := range field {
		pos = append(pos, Pair{num, p.pos.x})
		vel = append(vel, Pair{num, p.vel.x})
		acc = append(acc, Pair{num, p.acc.x})
	}
	sort.Slice(pos[:], func(i, j int) bool { return pos[i].val < pos[j].val })
	sort.Slice(vel[:], func(i, j int) bool { return vel[i].val < vel[j].val })
	sort.Slice(acc[:], func(i, j int) bool { return acc[i].val < acc[j].val })
	for n := 0; n < len(pos); n++ {
		if !loose_order_match(n, pos, vel) || !loose_order_match(n, pos, acc) {
			return false
		}
	}
	return true
}

func main() {
	field := make(map[int]Particle)

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	num := 0
	for scanner.Scan() {
		var p Particle
		_, err = fmt.Sscanf(scanner.Text(), "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>\n", &p.pos.x, &p.pos.y, &p.pos.z, &p.vel.x, &p.vel.y, &p.vel.z, &p.acc.x, &p.acc.y, &p.acc.z)
		if err != nil {
			panic("scanf error")
		}
		field[num] = p
		num++
	}

	for i := 0; true; i++ {
		collisions := make(map[Point][]int)

		// Update positions
		for num, p := range field {
			p.vel.x += p.acc.x
			p.vel.y += p.acc.y
			p.vel.z += p.acc.z

			p.pos.x += p.vel.x
			p.pos.y += p.vel.y
			p.pos.z += p.vel.z

			field[num] = p

			collisions[p.pos] = append(collisions[p.pos], num)
		}

		// Reap collisions
		for _, cols := range collisions {
			if len(cols) > 1 {
				for _, n := range cols {
					delete(field, n)
				}
			}
		}

		// Check if no more collisions can occur
		if particles_sorted(field) {
			fmt.Printf("Order stabilized at iteration %d\n", i)
			break
		}

	}

	fmt.Printf("Particles remaining: %d\n", len(field))
}
