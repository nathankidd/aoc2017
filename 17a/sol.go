package main

import "fmt"

const input = 328

func main() {
	l := []int{0}
	pos := 0
	for i := 1; i <= 2017; i++ {
		pos = (pos + input) % len(l)
		l = append(l, 0)
		copy(l[pos+1:], l[pos:])
		pos++
		l[pos] = i
	}

	fmt.Println(l[(pos+1)%len(l)])
}
