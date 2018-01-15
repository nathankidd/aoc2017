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
	// Example inputs
	//gena := newGenerator(65, 16807)
	//genb := newGenerator(8921, 48271)

	// Real inputs
	gena := newGenerator(618, 16807)
	genb := newGenerator(814, 48271)

	pairs := 0

	for i := 0; i < 5*1000*1000; i++ {
		var a uint64
		var b uint64
		for {
			a = gena()
			if a & 3 == 0 {
				break
			}
		}
		for {
			b = genb()
			if b & 7 == 0 {
				break
			}
		}
		if false && i < 5 {
			fmt.Printf("A=%d\t", a)
			fmt.Printf("B=%d\n", b)
		}
		a &= 0xFFFF
		b &= 0xFFFF
		if a == b {
			pairs++
		}

	}

	fmt.Printf("Pairs: %d\n", pairs)
}
