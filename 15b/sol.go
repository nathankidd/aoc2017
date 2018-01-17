package main

import "fmt"
import "math/bits"

func newGenerator(init, factor uint64, multiple uint) func() uint64 {
	n := init
	if bits.OnesCount(multiple) != 1 {
		panic("multiple is not a power of two")
	}
	mask := uint64(multiple - 1)
	return func() uint64 {
		for {
			n *= factor
			n %= 2147483647 // TODO &= 0x7fffffff something
			if n & mask == 0 {
				break
			}
		}
		return n
	}
}

func main() {
	gena := newGenerator(618, 16807, 4)
	genb := newGenerator(814, 48271, 8)

	pairs := 0

	for i := 0; i < 5*1000*1000; i++ {
		a := gena() & 0xffff
		b := genb() & 0xffff
		if a == b {
			pairs++
		}

	}

	fmt.Printf("Pairs: %d\n", pairs)
}
