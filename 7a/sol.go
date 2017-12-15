package main

import "fmt"
import "os"
import "bufio"
import "strings"

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tree := make(map[string]string)
	var root string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var parent string
		for pos, token := range strings.Fields(scanner.Text()) {
			if token[len(token)-1] == ',' {
				token = token[:len(token)-1]
			}
			switch pos {
			case 0:
				parent = token
				root = token // any starting point
			case 1: // ignore weight
			case 2: // ignore ->
			default: // child nodes
				tree[token] = parent
			}
		}
	}
	for {
		if _, exist := tree[root]; !exist {
			break
		}
		root = tree[root]
	}

	fmt.Printf("Root node: %s\n", root)
}
