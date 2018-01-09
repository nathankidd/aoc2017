// This is horribly slow
package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

type Layer struct {
	siz int // range of this layer
	dir int // direction of scanner
	pos int // current scanner position
}

func main() {
	fwall := make(map[int]Layer)
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
		fwall[depth] = Layer{size, 1, 0}
		if depth > max {
			max = depth
		}
	}

	for i := 0; i <= max; i++ {
		if _, exists := fwall[i]; !exists {
			fwall[i] = Layer{0, 0, 0}
		}
	}
	// Odd delays always catch us on the 2nd filter, so skip them
	for delay := 0; delay > -1000000; delay -= 2 {
		hit := false
		// reset positions
		for i := 0; i <= max; i++ {
			fwall[i] = Layer{fwall[i].siz, 1, 0}
		}
		if -delay%1000 == 0 {
			fmt.Printf("%d\n", -delay)
		}
		//fmt.Printf("---- %d ------------\n", delay)
		for pos := delay; pos <= max; pos++ {
			//fmt.Printf("[%2d->%2d] ", pos, fwall[pos].siz)
			for l := 0; l <= max; l++ {
				f := fwall[l]
				if f.siz == 0 {
					//			if pos == l {
					//			fmt.Printf("  .")
					//			} else {
					//			fmt.Printf("   ")
					//			}
					continue
				}
				if pos == l && f.pos == 0 {
					hit = true
				}
				//		if pos == l {
				//			if f.pos == 0 {
				//			fmt.Printf("%2d*", f.pos)
				//			} else {
				//			fmt.Printf("%2d.", f.pos)
				//			}
				//		} else {
				//		fmt.Printf("%2d ", f.pos)
				//		}
				f.pos += f.dir
				if f.pos == 0 || f.pos == f.siz-1 {
					f.dir = -f.dir
				}
				fwall[l] = f
			}
			//fmt.Printf("\n")
		}
		//fmt.Printf("Delay %d Severity: %d\n", -delay, severity)
		if !hit {
			fmt.Printf("Delay %d\n", -delay)
			break
		}
	}
}
