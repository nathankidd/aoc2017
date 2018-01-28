package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "errors"

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

func findPos(prog [proglen]rune, c rune) (int, error) {
	for i := 0; i < proglen; i++ {
		if prog[i] == c {
			return i, nil
		}
	}
	return -1, errors.New("can't find rune")
}

func main() {
	var prog [proglen]rune
	head := 0
	for i := 0; i < proglen; i++ {
		prog[i] = rune('a' + i)
	}

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var toks []string
	for scanner.Scan() {
		toks = strings.Split(scanner.Text(), ",")
	}
	for round := 0; round < 1; round++ {
		if round%100 == 0 {
			fmt.Printf("%d\n", round)
		}
		for _, tok := range toks {
			//for i := 0; i < proglen; i++ {
			//	fmt.Printf("%c", prog[(head+i)%proglen])
			//}
			//fmt.Printf("\ntok=%s ", tok)
			var sx int
			_, err := fmt.Sscanf(tok, "s%d", &sx)
			if err == nil {
				head = (head + (proglen - sx)) % proglen
				continue
			}
			var xa, xb int
			_, err = fmt.Sscanf(tok, "x%d/%d", &xa, &xb)
			if err == nil {
				xa = (head + xa) % proglen
				xb = (head + xb) % proglen
				prog[xa], prog[xb] = prog[xb], prog[xa]
				continue
			}
			var pa, pb rune
			_, err = fmt.Sscanf(tok, "p%c/%c", &pa, &pb)
			if err == nil {
				a, err := findPos(prog, pa)
				if err != nil {
					panic(err)
				}
				b, err := findPos(prog, pb)
				if err != nil {
					panic(err)
				}
				prog[a], prog[b] = prog[b], prog[a]
				continue
			}

			fmt.Printf("invalid input '%s'\n", tok)
			os.Exit(1)
		}
	}

	fmt.Printf("Program sequence: ")
	for i := 0; i < proglen; i++ {
		fmt.Printf("%c", prog[(head+i)%proglen])
	}
	fmt.Printf("\n")
}
