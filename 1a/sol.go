package main

import "fmt"
import "io/ioutil"

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	if input[len(input)-1] == '\n' {
		input = input[:len(input)-1]
	}
	for i,_ := range input {
		input[i] -= '0' //  atoi
	}
	prev := input[len(input)-1]
	sum := 0
	for _,cur  := range input {
		if prev == cur {
			sum += int(cur)
		}
		prev = cur
	}
	fmt.Printf("Sum: %d\n", sum);
}
