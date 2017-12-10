package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "math"

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	checksum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		min := math.MaxInt64 // TODO warn on overflow
		max := 0
		for _, token := range strings.Fields(scanner.Text()) {
			num, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}
		fmt.Printf("%d -> %d: %s\n", min, max, scanner.Text())
		checksum += max - min
	}
	fmt.Printf("Checksum: %d\n", checksum);
}
