package main

import "fmt"
import "os"
import "bufio"

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [201][]byte
	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		grid[line] = []byte(scanner.Text())
		line++
	}
	for y := 0; y < line; y++ {
		for x := 0; x < line; x++ {
			fmt.Printf("%c", grid[y][x])
		}
		fmt.Printf("\n")
	}

	// Find starting point
	x := 0
	y := 0
	for ; grid[y][x] == ' '; x++ {
	}

	// Follow path
	dx := 0
	dy := 1

	steps := 0
loop:
	for {
		switch grid[y][x] {
		case '+':
			// direction change
			if dx == 0 {
				dy = 0
				// was vertical, look left and right
				if grid[y][x-1] == '-' {
					dx = -1
				} else if grid[y][x+1] == '-' {
					dx = 1
				} else {
					break loop
				}
			} else {
				dx = 0
				// was horizontal , look up and down
				if grid[y-1][x] == '|' {
					dy = -1
				} else if grid[y+1][x] == '|' {
					dy = 1
				} else {
					break loop
				}
			}
		case ' ':
			break loop
		case '-':
		case '|':
		default:
			fmt.Printf("%c", grid[y][x])
		}

		x += dx
		y += dy
		steps++
	}
	fmt.Printf("\n")
	fmt.Printf("steps: %d\n", steps)

}
