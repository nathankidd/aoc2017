package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func main() {
	var wall []int
	hwall := make(map[int]int)

	maxfactor := 1
	max := 0

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ": ")
		if len(s) != 2 {
			fmt.Errorf("Invalid input: '%s'\n", scanner.Text())
			os.Exit(1)
		}
		depth, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}
		size, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		dist := size * 2 // unrolled travel distance
		dist -= 1        // zero-based
		dist -= 1        // don't repeat bottom position
		hwall[depth] = dist
		maxfactor *= dist
		fmt.Printf("hwall[%2d] = %d * 2 - 2 = %d\n", depth, size, hwall[depth])
		if depth > max {
			max = depth
		}
	}

	for i := 0; i <= max; i++ {
		wall = append(wall, hwall[i])
	}

	fmt.Printf("maxfactor: %d\n", maxfactor)
	for delay := 0; delay < maxfactor; delay++ {
		if delay%1000000 == 0 {
			fmt.Printf("-- %d\n", delay)
		}
		hit := false
		for pos := 0; !hit && pos <= max; pos++ {
			if wall[pos] == 0 {
				continue
			}
			if (pos+delay)%wall[pos] == 0 {
				hit = true
				break
			}
		}
		if !hit {
			fmt.Printf("Delay %d\n", delay)
			break
		}
	}
}
