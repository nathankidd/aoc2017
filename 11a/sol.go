package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "math"

func Abs(v int) int {
	if v >= 0 {
		return v
	}
	return -v
}

func DistanceTo(x, y float64) float64 {
	dist := 0.
	x = math.Abs(x)
	y = math.Abs(y)

	if y >= x {
		dist += x * 2. // diagonal steps, each whole number is two moves diagonal
		dist += y - x  // remaining vertical steps, whole number is one move 
	} else {
		// Cannot move directly on horizontal, must do extra vertical steps
		dist += y * 2. // diagonal steps to reach y
		dist += (x - y) * 2. // each remaining horizontal unit needs two moves 
	}
	return dist
}

func DistanceTo2(dirs map[string]int) int {
	// doesn't actually work
	delta := 0
	delta += Abs(dirs["n"] -  dirs["s"])
	delta += Abs(dirs["ne"] -  dirs["sw"])
	delta += Abs(dirs["nw"] -  dirs["se"])
	return delta
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	x := 0.0
	y := 0.0

	maxdist := 0.
	dirs := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, tok := range strings.Split(scanner.Text(), ",") {
			dirs[tok]++
			switch tok {
			case "n":
				y += 1.0
			case "ne":
				y += 0.5
				x += 0.5
			case "se":
				y -= 0.5
				x += 0.5
			case "s":
				y -= 1.0
			case "sw":
				y -= 0.5
				x -= 0.5

			case "nw":
				y += 0.5
				x -= 0.5
			default:
				fmt.Errorf("invalid input '%s'\n", tok)
				os.Exit(1)
			}
			dist := DistanceTo(x, y)
			if dist > maxdist {
				maxdist = dist
			}
		}
	}
	fmt.Printf("x=%.1f y=%.1f\n", x, y)
	fmt.Print(dirs)
	fmt.Printf("\n")

	// We don't care about direction, only distance, so normalize direction

	fmt.Printf("Distance: %f (max: %f)\n", DistanceTo(x, y), maxdist)

	fmt.Printf("Method2: %d\n", DistanceTo2(dirs))

}
