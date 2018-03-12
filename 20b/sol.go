package main

import "fmt"
import "os"
import "bufio"

func Abs(v int) int {
	if v >= 0 {
		return v
	}
	return -v
}

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

const MaxInt = int(^uint(0) >> 1)

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

	// TODO determine when no more particles can collide
	for i := 0; i < 1000; i++ {
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
	}

	fmt.Printf("Left: %d\n", len(field))
}
