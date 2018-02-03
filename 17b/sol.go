package main

import "fmt"

const input = 328

type Node struct {
	v    int
	next *Node
}

func newNode(v int) *Node {
	n := new(Node)
	n.v = v
	return n
}

func insertNode(parent *Node, v int) *Node {
	n := newNode(v)
	n.next = parent.next
	parent.next = n
	return n
}

//func printList(start *Node) {
//	fmt.Printf("[ %d", start.v)
//	for n := start.next; n != start; n = n.next {
//		fmt.Printf(" %d", n.v)
//	}
//	fmt.Printf(" ]\n")
//}

func main() {
	head := newNode(0)
	head.next = head

	pos := head
	for i := 1; i <= 50000000; i++ {
		if i%1000000 == 0 {
			fmt.Println(i)
		}
		//printList(head)
		for p := 0; p < input; p++ {
			pos = pos.next
		}
		pos = insertNode(pos, i)
	}

	n := head
	for ; n.v != 0; n = n.next {
	}
	fmt.Printf("Result: %d\n", n.next.v)
}
