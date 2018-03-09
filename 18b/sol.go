package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "time"

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
	xreg bool
	Y    int
	yreg bool
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

func RegOrImmediate(s string) (int, bool) {
	usereg := false
	v, err := strconv.Atoi(s)
	if err != nil {
		usereg = true
		v = regindex(s)
	}
	return v, usereg
}

func run(progid int, code []Op, in chan int, out chan int, ret chan int) {
	var logging = false
	var reg [p + 1]int
	reg[p] = progid
	numsnd := 0
	ip := 0
loop:
	for {
		nextip := ip + 1
		op := code[ip]
		var xval int
		if op.xreg {
			xval = reg[op.X]
		} else {
			xval = op.X
		}
		var yval int
		if op.yreg {
			yval = reg[op.Y]
		} else {
			yval = op.Y
		}
		if logging {
			fmt.Printf("[%2d] %s ", ip, opname(op.code))
			if op.xreg {
				fmt.Printf(" %s (%d),", regname(op.X), xval)
			} else {
				fmt.Printf(" %d,", xval)
			}
			if op.yreg {
				fmt.Printf(" %s (%d)", regname(op.Y), yval)
			} else {
				fmt.Printf(" %d", yval)
			}
			fmt.Printf("\n")
		}
		switch op.code {
		case snd:
			select {
			case out <- xval:
				numsnd++
			case <-time.After(5 * time.Second):
				break loop
			}
		case set:
			reg[op.X] = yval
		case add:
			reg[op.X] += yval
		case mul:
			reg[op.X] *= yval
		case mod:
			reg[op.X] %= yval
		case rcv:
			select {
			case reg[op.X] = <-in:
			case <-time.After(5 * time.Second):
				break loop
			}
		case jgz:
			if xval > 0 {
				nextip = ip + yval
			}
		}
		ip = nextip
		if ip < 0 || ip >= len(code) {
			fmt.Printf("Instructions escape code\n")
			os.Exit(1)
		}
	}
	ret <- numsnd
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var code []Op
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		toks := strings.Fields(scanner.Text())
		var op int
		X, xreg := RegOrImmediate(toks[1])
		Y, yreg := 0, false
		if len(toks) == 3 {
			Y, yreg = RegOrImmediate(toks[2])
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
		code = append(code, Op{op, X, xreg, Y, yreg})
	}

	ch0 := make(chan int, 100)
	ch1 := make(chan int, 100)
	ret0 := make(chan int)
	ret1 := make(chan int)
	go run(0, code, ch0, ch1, ret0)
	go run(1, code, ch1, ch0, ret1)

	fmt.Printf("Prog 1 sent %d times\n", <-ret1)

}
