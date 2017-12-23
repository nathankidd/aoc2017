package main

import "fmt"
import "os"
import "bufio"
import "encoding/hex"

func CalcDenseHash(input *[256]byte) *[16]byte {
	var dhash = new([16]byte)
	for i := 0; i < len(input); i+= 16 {
		out := input[i]
		for j := 1; j < 16; j++ {
			out ^= input[i+j]
		}
		dhash[i>>4] = out
	}
	return dhash
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lengths []byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lengths = []byte(scanner.Text())
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)

	var list[256]byte;
	for i := 0; i < len(list); i++ {
		list[i] = byte(i)
	}
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
	dhash := CalcDenseHash(&list)

	fmt.Printf("Result: %s\n", hex.EncodeToString(dhash[:]))
}
