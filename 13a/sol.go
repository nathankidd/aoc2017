package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

type Layer struct {
	siz  int // range of this layer
	dir  int // direction of scanner
	pos  int // current scanner position
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
			fwall[i] = Layer{0,0,0}
		}
	}
	severity := 0
	for pos := 0; pos <= max; pos++ {
		fmt.Printf("[%2d->%2d] ", pos, fwall[pos].siz)
		for l := 0; l <= max; l++ {
			f := fwall[l]
			if f.siz == 0 {
				fmt.Printf("   ")
				continue
			}
			if pos == l && f.pos == 0 {
				severity += pos * f.siz
			}
			if pos == l {
				if f.pos == 0 {
					fmt.Printf("%2d*", f.pos)
				} else {
					fmt.Printf("%2d.", f.pos)
				}
			} else {
				fmt.Printf("%2d ", f.pos)
			}
			f.pos += f.dir
			if f.pos == 0 || f.pos == f.siz-1 {
				f.dir = -f.dir
			}
			fwall[l] = f
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Severity: %d\n", severity)
}
