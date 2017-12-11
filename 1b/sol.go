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
	for i, _ := range input {
		input[i] -= '0' //  atoi
	}
	size := len(input)
	sum := 0
	for pos, cur := range input {
		oppo := input[(pos+size/2)%size]
		if oppo == cur {
			sum += int(cur)
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}
