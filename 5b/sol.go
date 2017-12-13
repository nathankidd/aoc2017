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

	var code []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, token := range strings.Fields(scanner.Text()) {
			jump, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			code = append(code, jump)
			fmt.Printf("%d\n", jump)
		}
	}
	fmt.Printf("Num jumps: %d\n", len(code))

	ip := 0
	instructions := 0
	for {
		jump := code[ip]
		if jump >= 3 {
			code[ip]--
		} else {
			code[ip]++
		}
		ip += jump
		instructions++
		if ip < 0 || ip >= len(code) {
			break
		}
	}
	fmt.Printf("Number of instructions to escape: %d\n", instructions)
}
