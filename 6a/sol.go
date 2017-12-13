package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

const numbanks = 16
type ChunkCtr uint32

func bankKey(bank [numbanks]ChunkCtr) string {
	var b []byte
	for _, n := range bank {
		b = append(b,
			byte(n>>24),
			byte((n>>16)&0xff),
			byte((n>>8)&0xff),
			byte(n&0xff))
	}
	return string(b);
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var bank [numbanks]ChunkCtr
	states := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for pos, token := range strings.Fields(scanner.Text()) {
			num, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			bank[pos] = ChunkCtr(num)
		}
	}

	iter := 0
	looplen := 0
	for {
		key := bankKey(bank)
		if states[key] != 0 {
			looplen = iter - states[key]
			break
		}
		states[key] = iter
		iter++

		var chunks ChunkCtr = 0
		idx := 0
		for i, size := range bank {
			if size > chunks {
				chunks = size
				idx = i
			}
		}
		for bank[idx] = 0; chunks > 0; chunks-- {
			idx = (idx + 1) % len(bank)
			bank[idx]++
		}
	}
	fmt.Printf("Iterations: %d\n", iter)
	fmt.Printf("Looplen: %d\n", looplen)
}
