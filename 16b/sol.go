package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "time"

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

type Op struct {
	code rune
	a, b ProgType
}

const proglen = 16

type ProgType uint8

func main() {
	var prog [proglen]ProgType
	var head ProgType
	for i := 0; i < proglen; i++ {
		prog[i] = ProgType('a' + i)
	}
	var ops []Op

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for pos, tok := range strings.Split(scanner.Text(), ",") {
			//for i := 0; i < proglen; i++ {
			//	fmt.Printf("%c", prog[(head+i)%proglen])
			//}
			//fmt.Printf("\ntok=%s ", tok)
			if pos < 10 {
				//			fmt.Printf(".%d %s\n", pos, tok)
			}
			var a, b int
			_, err := fmt.Sscanf(tok, "s%d", &a)
			if err == nil {
				ops = append(ops, Op{'s', ProgType(a), ProgType(b)})
				continue
			}
			_, err = fmt.Sscanf(tok, "x%d/%d", &a, &b)
			if err == nil {
				ops = append(ops, Op{'x', ProgType(a), ProgType(b)})
				continue
			}
			var pa, pb rune
			_, err = fmt.Sscanf(tok, "p%c/%c", &pa, &pb)
			if err == nil {
				ops = append(ops, Op{'p', ProgType(pa), ProgType(pb)})
				continue
			}

			fmt.Printf("invalid input '%s'\n", tok)
			os.Exit(1)
		}
	}

	//	for i := 0; i < 10; i++ {
	//		fmt.Printf("#%d %c%c/%c\n", i, ops[i].code, rune(ops[i].a), rune(ops[i].b))
	//	}

	start := time.Now()
	for round := 0; round < 1*1000*1000*1000; round++ {
		if round%1000000 == 0 {
			fmt.Printf("%d\n", round)
		}
		for _, op := range ops {
			switch op.code {
			case 's':
				head = (head + (proglen - op.a)) & (proglen - 1)
			case 'x':
				a := (head + op.a) & (proglen - 1)
				b := (head + op.b) & (proglen - 1)
				prog[a], prog[b] = prog[b], prog[a]
			case 'p':
				a := 0
				for ; prog[a] != (op.a); a++ {
				}
				b := 0
				for ; prog[b] != (op.b); b++ {
				}
				prog[a], prog[b] = prog[b], prog[a]
			default:
				fmt.Printf("invalid op '%c'\n", op.code)
				os.Exit(1)
			}
		}
	}
	end := time.Now()

	fmt.Printf("%dms\n", end.Sub(start)/1000/1000)

	fmt.Printf("Program sequence: ")
	for i := ProgType(0); i < proglen; i++ {
		fmt.Printf("%c", prog[(head+i)%proglen])
	}
	fmt.Printf("\n")
}
