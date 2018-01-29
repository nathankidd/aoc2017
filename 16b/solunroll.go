// Unrolled binary is actually slower than sol.go (1900ms vs 1800ms for 10K loops)
package main

import "fmt"
import "os"
import "bufio"
import "strings"

// There are sixteen programs in total, named a through p. They start by standing in a line: a stands in position 0, b stands in position 1, and so on until p, which stands in position 15.
//
// The programs' dance consists of a sequence of dance moves:
//
// Spin, written sX, makes X programs move from the end to the front, but maintain their order otherwise. (For example, s3 on abcde produces cdeab).
// Exchange, written xA/B, makes the programs at positions A and B swap places.
// Partner, written pA/B, makes the programs named A and B swap places.
//
// For example, with only five programs standing in a line (abcde), they could do the following dance:
//
// 	s1, a spin of size 1: eabcd.
// 	x3/4, swapping the last two programs: eabdc.
// 	pe/b, swapping programs e and b: baedc.

const proglen = 16
const progmask = proglen - 1

type ProgType uint8

func main() {
	fmt.Printf(`
// DO NOT EDIT -- Produced by solunroll.go
package main
import "fmt"
import "time"
const proglen = 16
const progmask = proglen - 1
type ProgType uint8
func main() {
	var prog [proglen]ProgType
	var head ProgType
	for i := 0; i < proglen; i++ {
		prog[i] = ProgType('a' + i)
	}
	start := time.Now()
	for round := 0; round < 1*1000*1000*1000; round++ {
		if round%%1000000 == 0 {
			fmt.Printf("%%d\n", round)
		}
		var a, b ProgType
`)

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, tok := range strings.Split(scanner.Text(), ",") {
			var a, b int
			_, err := fmt.Sscanf(tok, "s%d", &a)
			if err == nil {
				fmt.Printf(`
				head = (head + (proglen - %d)) & progmask
				`, a)
				continue
			}
			_, err = fmt.Sscanf(tok, "x%d/%d", &a, &b)
			if err == nil {
				fmt.Printf(`
				a = (head + %d) & progmask
				b = (head + %d) & progmask
				prog[a], prog[b] = prog[b], prog[a]
				`, a, b)
				continue
			}
			var pa, pb rune
			_, err = fmt.Sscanf(tok, "p%c/%c", &pa, &pb)
			if err == nil {
				fmt.Printf(`
				a = 0
				for ; prog[a] != (%d); a++ {
				}
				b = 0
				for ; prog[b] != (%d); b++ {
				}
				prog[a], prog[b] = prog[b], prog[a]
				`, pa, pb)
				continue
			}

			fmt.Printf("invalid input '%s'\n", tok)
			os.Exit(1)
		}
	}

	fmt.Printf(`
	}
	end := time.Now()

	fmt.Printf("%%dms\n", end.Sub(start)/1000/1000)

	fmt.Printf("Program sequence: ")
	for i := ProgType(0); i < proglen; i++ {
		fmt.Printf("%%c", prog[(head+i)%%proglen])
	}
	fmt.Printf("\n")
}
	`)
}
