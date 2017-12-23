package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lengths []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, tok := range strings.Split(scanner.Text(), ",") {
			num, err := strconv.Atoi(tok)
			if err != nil {
				panic(err)
			}
			if num > 255 {
				fmt.Errorf("Invalid input: '%s'\n", scanner.Text())
				os.Exit(1)
			}
			lengths = append(lengths, num)
			fmt.Printf("%d,", num)
		}
	}
	fmt.Printf("\n")

	var list[256]int;
	for i := 0; i < len(list); i++ {
		list[i] = i
	}
	pos := 0
	skip := 0
	for _, length := range lengths {
		for start, end := pos, pos + length - 1; start < end; start, end = start + 1, end - 1 {
			list[start%len(list)], list[end%len(list)] = list[end%len(list)], list[start%len(list)]
		}
		pos = (pos + length + skip) % len(list)
		skip++
	}

	fmt.Printf("Result: %d\n", list[0] * list[1])
}
