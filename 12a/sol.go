package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func GetGroup(node int, gmap *map[int][]int, group *map[int]bool) {
	for _, n := range (*gmap)[node] {
		if _, exist := (*group)[n]; !exist {
			(*group)[n] = true
			GetGroup(n, gmap, group)
		}
	}
}

func main() {
	gmap := make(map[int][]int)

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		node := 0
		for pos, tok := range strings.Fields(scanner.Text()) {
			if pos == 1 {
				continue // skip <->
			}
			if tok[len(tok)-1] == ',' {
				tok = tok[:len(tok)-1]
			}
			num, err := strconv.Atoi(tok)
			if err != nil {
				panic(err)
			}
			switch pos {
			case 0:
				node = num
			default:
				gmap[node] = append(gmap[node], num)
			}
		}
	}

	group := map[int]bool {
		0: true,
	}
	GetGroup(0, &gmap, &group)
	fmt.Printf("Group 0 members: %d\n", len(group))
}
