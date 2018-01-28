package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "time"
import "math/bits"

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

type ProgType uint64

type Op struct {
	code rune
	a, b uint64
	mask uint64
}

const proglen = 16

func printProg(prog uint64) {
	for i := uint(0); i < proglen; i++ {
		fmt.Printf("%c", ((prog>>(i*4))&0xf)+'a')
	}
}

func main() {
	// 16 4-bit values stored in one 64-bit value
	// 4-bit values are 0-based, not actual ascii values
	var prog uint64
	for i := uint64(0); i < proglen; i++ {
		prog |= i << (i * 4)
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
			if pos < 0 {
				fmt.Printf(".%d %s\n", pos, tok)
			}
			var a, b uint64
			_, err := fmt.Sscanf(tok, "s%d", &a)
			if err == nil {
				ops = append(ops, Op{'s', a, b, 0})
				continue
			}
			_, err = fmt.Sscanf(tok, "x%d/%d", &a, &b)
			if err == nil {
				ops = append(ops, Op{'x', a, b, ^(0xf<<(a*4) | 0xf<<(b*4))})
				continue
			}
			var pa, pb rune
			_, err = fmt.Sscanf(tok, "p%c/%c", &pa, &pb)
			if err == nil {
				a = uint64(pa - 'a')
				b = uint64(pb - 'a')
				ops = append(ops, Op{'p', a, b, 0})
				continue
			}

			fmt.Printf("invalid input '%s'\n", tok)
			os.Exit(1)
		}
	}

	for i := 0; i < 0; i++ {
		if ops[i].code == 'p' {
			fmt.Printf("#%d %c%c/%c\n", i, ops[i].code, rune(ops[i].a)+'a', rune(ops[i].b)+'a')
		} else {
			fmt.Printf("#%d %c%d/%d\n", i, ops[i].code, rune(ops[i].a), rune(ops[i].b))
		}
	}

	start := time.Now()
	for round := 0; round < 10000; round++ {
		//for round := 0; round < 1*1000*1000*1000; round++ {
		//		if round%1000000 == 0 {
		//			fmt.Printf("%d\n", round)
		//		}
		for _, op := range ops {
			//printProg(prog)
			switch op.code {
			case 's':
				//	fmt.Printf(" %c%d\n", op.code, op.a)
				prog = bits.RotateLeft64(prog, int(op.a*4))
			case 'x':
				//	fmt.Printf(" %c%d/%d\n", op.code, op.a, op.b)
				a := op.a
				b := op.b
				aval := (prog >> (a * 4)) & 0xf
				bval := (prog >> (b * 4)) & 0xf
				//mask := uint64(^(0xf<<(a*4) | 0xf<<(b*4)))
				prog &= op.mask
				prog |= aval << (b * 4)
				prog |= bval << (a * 4)
			case 'p':
				//	fmt.Printf(" %c%c/%c", op.code, op.a+'a', op.b+'a')
				a := uint(0)
				for v := prog; v&0xf != op.a; a++ {
					v >>= 4
				}
				b := uint(0)
				for v := prog; v&0xf != op.b; b++ {
					v >>= 4
				}
				aval := (prog >> (a * 4)) & 0xf
				bval := (prog >> (b * 4)) & 0xf
				mask := uint64(^(0xf<<(a*4) | 0xf<<(b*4)))
				//fmt.Printf(" [%d]=%d [%d]=%d k=%x\n", a, aval, b, bval, mask)
				prog &= mask
				prog |= aval << (b * 4)
				prog |= bval << (a * 4)
			default:
				fmt.Printf("invalid op '%c'\n", op.code)
				os.Exit(1)
			}
		}
	}
	end := time.Now()

	fmt.Printf("%dms\n", end.Sub(start)/1000/1000)

	fmt.Printf("Program sequence: ")
	printProg(prog)
	fmt.Printf("\n")
}
