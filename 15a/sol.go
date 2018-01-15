package main

import "fmt"

func newGenerator(init, factor uint64) func() uint64 {
	n := uint64(init)
	return func() uint64 {
		n *= factor
		n %= 2147483647 // TODO &= 0x7fffffff something
		return n
	}
}

func main() {
	// Generator A starts with 618, factor 16807
	// Generator B starts with 814, factor 48271

	// These generate 1, for first 5 pairs
	//gena := newGenerator(65, 16807)
	//genb := newGenerator(8921, 48271)

	gena := newGenerator(618, 16807)
	genb := newGenerator(814, 48271)

	pairs := 0

	for i := 0; i < 40*1000*1000; i++ {
		a := gena() & 0xFFFF
		b := genb() & 0xFFFF
		if a == b {
			pairs++
		}
		if false && i < 5 {
			fmt.Printf("A=%d\t", a)
			fmt.Printf("B=%d\n", b)
		}

	}

	fmt.Printf("Pairs: %d\n", pairs)
}
