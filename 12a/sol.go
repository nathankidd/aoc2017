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

func GroupsEqual(a, b map[int]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for key, _ := range a {
		if _, exists := b[key]; !exists {
			return false
		}
	}
	return true
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

	ggmap := make(map[int]map[int]bool)
	for key, _ := range gmap {
		g := make(map[int]bool)
		GetGroup(key, &gmap, &g)
		ggmap[key] = g
	}

	used := make(map[int]bool)
	groupcnt := 0
	for key, _ := range gmap {
		if used[key] == true {
			continue
		}
		for testkey, _ := range gmap {
			a := ggmap[key]
			b := ggmap[testkey]
			if GroupsEqual(a, b) {//ggmap[key], ggmap[testkey]) {
				used[testkey] = true
			}
		}
		groupcnt++
//		fmt.Printf("group %d\n", key)
	}
	fmt.Printf("Separate groups: %d\n", groupcnt)
}
