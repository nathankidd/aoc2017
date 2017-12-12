package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "sort"

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	numvalid := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := make(map[string]int)
		valid := true
		for _, token := range strings.Fields(scanner.Text()) {
			r := []rune(token)
			sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
			token = string(r)
			words[token]++
			if words[token] > 1 {
				valid = false
			}
		}
		if valid {
			numvalid++
		}
	}

	fmt.Printf("Num valid: %d\n", numvalid)
}
