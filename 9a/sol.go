package main

import "fmt"
import "os"
import "bufio"

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = []rune(scanner.Text())
	}

	garbagechars := 0
	garbage := false
	level := 0
	score := 0
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '!':
			i++
			continue
		case '>':
			if !garbage {
				fmt.Errorf("unmatched garbage end\n")
			}
			garbage = false
		default:
			if garbage {
				garbagechars++
				continue
			}
			switch input[i] {
			case '<':
				garbage = true
			case '{':
				level++
				score += level
			case '}':
				if level < 1 {
					fmt.Errorf("unmatched group end\n")
				}
				level--
			default:
			}
		}
	}

	fmt.Printf("Score: %d\n", score)
	fmt.Printf("Garbage chars: %d\n", garbagechars)
}
