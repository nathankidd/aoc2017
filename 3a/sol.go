package main

import "fmt"

const dest = 325489

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func main() {

	// Spiraling counter-clockwise means square ends in lower-right corner.
	// Square root adds 2 each winding, therefore the lower right corners
	// increase like 1^2, 3^3, 5^5, 7^7 ...

	sqroot := 1
	for {
		if sqroot * sqroot >= dest {
			break
		}
		sqroot += 2
	}
	delta := sqroot * sqroot - dest
	fmt.Printf("delta: %d (%d^2) - %d =  %d\n", sqroot*sqroot, sqroot, dest, delta)

	// Distance from outside edge to center
	halfsq := sqroot / 2
	// Offset from 1st block along an outside edge (don't care which)
	offset := delta % (sqroot - 1)
	// Convert to offset from center
	offset -= halfsq

	// We don't care which are x or y since we only need total distance
	fmt.Printf("halfsq = %d, offset = %d\n", halfsq, offset)
	fmt.Printf("Distance: %d\n", abs(halfsq)+abs(offset))
}
