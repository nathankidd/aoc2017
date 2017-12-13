// TODO This solution is very suboptimal
// There is no need to keep a 2D array, or track the x/y co-ordinates.
// Instead, track each winding in an array, and save the previous winding.
// Iterate the arrays linearly, tweaking behaviour when on a "corner".
package main

import "fmt"
import "os"

const dest = 325489

const center = 10 // "10 layers should be enough for anyone"
const arrsz = center*2 + 1

func getsum(grid [arrsz][arrsz]int, x int, y int) int {
	sum := 0
	sum += grid[x+1][y+0]
	sum += grid[x+1][y+1]
	sum += grid[x+0][y+1]
	sum += grid[x-1][y+1]
	sum += grid[x-1][y+0]
	sum += grid[x-1][y-1]
	sum += grid[x+0][y-1]
	sum += grid[x+1][y-1]
	fmt.Printf("%2d, %2d = %d\n", x-center, y-center, sum)
	return sum
}

func main() {

	var grid [arrsz][arrsz]int

	x := center
	y := center
	grid[x][y] = 1
	sqroot := 1
	found := 0
	for {
		fmt.Printf("------------------------- layer %d\n", sqroot)
		sqroot += 2
		sidelen := sqroot - 1
		if sidelen > center {
			fmt.Errorf("array too small")
			break
		}

		// Reposition to lower-right corner of new layer
		x++
		y--

		// Right side
		for i := 0; i < sidelen; i++ {
			y++
			grid[x][y] = getsum(grid, x, y)
			if grid[x][y] >= dest {
				found = grid[x][y]
				break
			}
		}
		if found != 0 {
			break
		}

		// Top side
		for i := 0; i < sidelen; i++ {
			x--
			grid[x][y] = getsum(grid, x, y)
			if grid[x][y] >= dest {
				found = grid[x][y]
				break
			}
		}
		if found != 0 {
			break
		}

		// Left side
		for i := 0; i < sidelen; i++ {
			y--
			grid[x][y] = getsum(grid, x, y)
			if grid[x][y] >= dest {
				found = grid[x][y]
				break
			}
		}
		if found != 0 {
			break
		}

		// Bottom side
		for i := 0; i < sidelen; i++ {
			x++
			grid[x][y] = getsum(grid, x, y)
			if grid[x][y] >= dest {
				found = grid[x][y]
				break
			}
		}
		if found != 0 {
			break
		}
	}
	if found == 0 {
		fmt.Errorf("something went wrong\n")
		os.Exit(1)
	}
	fmt.Printf("sqroot=%d, x = %d, y = %d\n", sqroot, x, y)
	fmt.Printf("Result: %d\n", found)
}
