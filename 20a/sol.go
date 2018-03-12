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

type Particle struct {
	particle int
	accel    int
}

const MaxInt = int(^uint(0) >> 1)

func main() {
	p := Particle{0, MaxInt}

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	particle := 0
	for scanner.Scan() {
		var x, y, z, vx, vy, vz, ax, ay, az int
		_, err = fmt.Sscanf(scanner.Text(), "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>\n", &x, &y, &z, &vx, &vy, &vz, &ax, &ay, &az)
		if err != nil {
			panic("scanf error")
		}
		accel := Abs(ax) + Abs(ay) + Abs(az)
		if accel < p.accel {
			p.accel = accel
			p.particle = particle
		}
		particle++
	}

	fmt.Printf("Particle %d has accel %d\n", p.particle, p.accel)
}
