package main

import "fmt"
import "math/bits"

func CalcHash(lengths []byte) *[16]byte {
	var list[256]byte;
	for i := 0; i < len(list); i++ {
		list[i] = byte(i)
	}
	lengths = append(lengths, []byte{17, 31, 73, 47, 23}...)
	pos := 0
	skip := 0
	for round := 0; round < 64; round++ {
		for _, length := range lengths {
			for start, end := pos, pos + int(length) - 1; start < end; start, end = start + 1, end - 1 {
				list[start%len(list)], list[end%len(list)] = list[end%len(list)], list[start%len(list)]
			}
			pos = (pos + int(length) + skip) % len(list)
			skip++
		}
	}
	var hash = new([16]byte)
	for i := 0; i < len(list); i+= 16 {
		out := list[i]
		for j := 1; j < 16; j++ {
			out ^= list[i+j]
		}
		hash[i>>4] = out
	}
	return hash
}

func main() {
	input := "ffayrhll"
	//input := "flqrgnkx" // test, output 8108

	blocks := 0

	for i := 0; i < 128; i++ {
		key := fmt.Sprintf("%s-%d", input, i)
		hash := CalcHash([]byte(key))
		fmt.Printf("%s = ", key)
		hashblocks := 0
		for pos, b := range hash {
			if false && i < 8 && pos < 1 {
				for bit := byte(1 << 7); bit > 0; bit >>= 1 {
					if (b & bit) == 0 {
						fmt.Printf(".")
					} else {
						fmt.Printf("#")
					}
				}
			}
			hashblocks += bits.OnesCount8(b)
		}
		fmt.Printf(" (%d)\n", hashblocks)
		blocks += hashblocks
	}

	fmt.Printf("Used blocks: %d\n", blocks)
}
