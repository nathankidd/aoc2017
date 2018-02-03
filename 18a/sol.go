package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

const (
	snd = iota
	set
	add
	mul
	mod
	rcv
	jgz
)
const (
	one = iota
	a
	b
	f
	i
	p
)

type Op struct {
	code int
	X    int
	Y    int
	reg  bool
}

func opname(o int) string {
	switch o {
	case snd:
		return "snd"
	case set:
		return "set"
	case add:
		return "add"
	case mul:
		return "mul"
	case mod:
		return "mod"
	case rcv:
		return "rcv"
	case jgz:
		return "jgz"
	default:
		panic("unknown reg")
	}
}
func regname(r int) string {
	switch r {
	case one:
		return "one"
	case a:
		return "a"
	case b:
		return "b"
	case f:
		return "f"
	case i:
		return "i"
	case p:
		return "p"
	default:
		panic("unknown reg")
	}
}

func regindex(r string) int {
	var X int
	switch r {
	case "1":
		X = one
	case "a":
		X = a
	case "b":
		X = b
	case "f":
		X = f
	case "i":
		X = i
	case "p":
		X = p
	default:
		panic("unknown reg")
	}
	return X
}
func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var code []Op
	var reg [p + 1]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		toks := strings.Fields(scanner.Text())
		var op int
		var X int
		var Y int
		usereg := false
		X = regindex(toks[1])
		if len(toks) == 3 {
			Y, err = strconv.Atoi(toks[2])
			if err != nil {
				usereg = true
				Y = regindex(toks[2])
			}
		}
		switch toks[0] {
		case "snd":
			op = snd
		case "set":
			op = set
		case "add":
			op = add
		case "mul":
			op = mul
		case "mod":
			op = mod
		case "rcv":
			op = rcv
		case "jgz":
			op = jgz
		default:
			panic("unknown opcode")
		}
		code = append(code, Op{op, X, Y, usereg})
	}

	rcvreg := 0
	ip := 0
loop:
	for {
		nextip := ip + 1
		op := code[ip]
		var val int
		if op.reg {
			val = reg[op.Y]
		} else {
			val = op.Y
		}
		fmt.Printf("[%2d] %s %s", ip, opname(op.code), regname(op.X))
		if op.reg {
			fmt.Printf(" %s (%d)\n", regname(op.Y), val)
		} else {
			fmt.Printf(" %d\n", val)
		}
		switch op.code {
		case snd:
			rcvreg = reg[op.X]
		case set:
			reg[op.X] = val
		case add:
			reg[op.X] += val
		case mul:
			reg[op.X] *= val
		case mod:
			reg[op.X] %= val
		case rcv:
			if op.X != 0 {
				fmt.Printf("First rcv: %d\n", rcvreg)
				break loop
			}
		case jgz:
			if reg[op.X] > 0 {
				nextip = ip + val
			}
		}
		ip = nextip
		if ip < 0 || ip >= len(code) {
			fmt.Printf("Instructions escape code\n")
			os.Exit(1)
		}
	}
}
