package main

import "fmt"

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

func Fill(g *[128][128]int, grp, x, y int) {
	if g[x][y] != -1 {
		return
	}
	g[x][y] = grp
	if x > 0 {
		Fill(g, grp, x - 1, y)
	}
	if y > 0 {
		Fill(g, grp, x, y - 1)
	}
	if x < 127 {
		Fill(g, grp, x + 1, y)
	}
	if y < 127 {
		Fill(g, grp, x, y + 1)
	}
}

func main() {
	input := "ffayrhll"

	var g [128][128]int

	for y := 0; y < 128; y++ {
		key := fmt.Sprintf("%s-%d", input, y)
		hash := CalcHash([]byte(key))
		fmt.Printf("%s = ", key)

		x := 0
		for _, b := range hash {
			for bit := byte(1 << 7); bit != 0; bit >>= 1 {
				if b & bit != 0 {
					g[x][y] = -1
				}
				x++
			}
		}
	}


	// Print usage map
	for y:= 0; y < 128; y++ {
		for x:= 0; x < 128; x++ {
			if g[x][y] < 0 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}

	grp := 0
	for y:= 0; y < 128; y++ {
		for x:= 0; x < 128; x++ {
			if g[x][y] < 0 { // used but ungrouped
				grp++
				Fill(&g, grp, x, y)
			}
		}
	}

	// Print partial group map
	for y:= 0; y < 128; y++ {
		for x:= 0; x < 128; x++ {
			if g[x][y] > 0 && g[x][y] < 10 {
				fmt.Printf("%d", g[x][y])
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Groups: %d\n", grp)
}
