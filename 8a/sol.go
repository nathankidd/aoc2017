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

	reg := make(map[string]int)
	var alltimemax_reg string
	var alltimemax_val int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tok := strings.Fields(scanner.Text())
		// b inc 5 if a > 1
		if len(tok) != 7 {
			fmt.Errorf("Invalid input: '%s'\n", scanner.Text())
			os.Exit(1)
		}
		regnum, err := strconv.Atoi(tok[2])
		if err != nil {
			panic(err)
		}
		ifnum, err := strconv.Atoi(tok[6])
		if err != nil {
			panic(err)
		}
		if _, exist := reg[tok[0]]; !exist {
			reg[tok[0]] = 0
		}
		if _, exist := reg[tok[4]]; !exist {
			reg[tok[4]] = 0
		}
		execute := false
		switch tok[5] {
		case "==":
			if reg[tok[4]] == ifnum {
				execute = true
			}
		case "!=":
			if reg[tok[4]] != ifnum {
				execute = true
			}
		case ">=":
			if reg[tok[4]] >= ifnum {
				execute = true
			}
		case "<=":
			if reg[tok[4]] <= ifnum {
				execute = true
			}
		case ">":
			if reg[tok[4]] > ifnum {
				execute = true
			}
		case "<":
			if reg[tok[4]] < ifnum {
				execute = true
			}
		default:
			fmt.Errorf("Invalid conditional operation: '%s'\n", scanner.Text())
			os.Exit(1)
		}
		if execute {
			switch tok[1] {
			case "inc":
				reg[tok[0]] += regnum
			case "dec":
				reg[tok[0]] -= regnum
			default:
				fmt.Errorf("Invalid register operation: '%s'\n", scanner.Text())
				os.Exit(1)
			}
		}
		if alltimemax_reg == "" || alltimemax_val < reg[tok[0]] {
			alltimemax_reg = tok[0]
			alltimemax_val = reg[tok[0]]
		}

	}

	var max string
	for key, value := range reg {
		if max == "" || reg[max] < value {
			max = key
		}
	}
	fmt.Printf("Max reg: %s = %d\n", max, reg[max])
	fmt.Printf("All time max reg: %s = %d\n", alltimemax_reg, alltimemax_val)
}
