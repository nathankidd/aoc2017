// DO NOT EDIT -- Produced by solunroll.go
package main

import "fmt"
import "time"

const proglen = 16
const progmask = proglen - 1

type ProgType uint8

func main() {
	var prog [proglen]ProgType
	var head ProgType
	for i := 0; i < proglen; i++ {
		prog[i] = ProgType('a' + i)
	}
	start := time.Now()
	for round := 0; round < 10000; round++ {
		if round%1000000 == 0 {
			fmt.Printf("%d\n", round)
		}
		var a, b ProgType

		head = (head + (proglen - 2)) & progmask

		a = (head + 5) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 0) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 14) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 3) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 14) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 0) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 10) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 13) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 15) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 7) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 3) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 12) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 7) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 4) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 8) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 15) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 4) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 5) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 13) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 2) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 6) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 14) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 7) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 4) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 6) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 14) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 14) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 1) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 15) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 5) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 14) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 12) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 3) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 7) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 5) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 14) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 0) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 4) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 15) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 7) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 15) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 9) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 11) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 9) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 1) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 7) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 0) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 11) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 15) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 4) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 12) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 6) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 7) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 12) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 0) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 0) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 9) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 3) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 7) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 1) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 7) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 10) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 13) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 7) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 8) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 0) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 11) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 14) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 5) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 8) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 7) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 11) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 11) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 1) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 1) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 2) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 13) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 6) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 12) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 6) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 5) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 15) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 1) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 7) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 8) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 6) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 3) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 0) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 12) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 6) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 7) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 6) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 14) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 7) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 10) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 6) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 6) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 11) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 13) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 7) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 9) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 14) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 7) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 15) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 5) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 0) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 11) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 15) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 1) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 11) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 9) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 7) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 13) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 1) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 10) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 12) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 1) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 3) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 5) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 6) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 10) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 9) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 3) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 7) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 15) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 7) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 11) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 14) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 6) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 3) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 6) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 7) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 6) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 8) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 14) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 11) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 1) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 6) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 15) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 3) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 3) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 15) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 3) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 12) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 13) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 14) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 14) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 0) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 7) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 13) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 15) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 7) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 9) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 8) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 7) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 9) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 9) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 11) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 14) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 14) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 6) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 11) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 12) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 6) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 4) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 14) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 12) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 6) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 12) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 6) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 0) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 4) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 15) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 9) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 12) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 6) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 11) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 6) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 15) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 14) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 9) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 8) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 11) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 8) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 7) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 6) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 8) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 6) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 13) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 0) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 15) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 0) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 15) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 1) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 0) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 11) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 15) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 11) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 8) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 5) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 0) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 10) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 2) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 3) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 0) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 3) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 0) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 11) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 14) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 6) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 2) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 7) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 6) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 6) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 2) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 0) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 6) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 6) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 4) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 13) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 6) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 7) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 6) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 1) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 0) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 0) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 4) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 6) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 7) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 3) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 11) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 1) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 15) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 7) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 11) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 14) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 7) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 8) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 1) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 1) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 12) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 6) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 2) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 2) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 14) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 8) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 8) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 3) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 7) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 4) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 8) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 12) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 2) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 14) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 2) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 0) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 5) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 11) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 6) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 14) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 6) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 2) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 14) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 1) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 3) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 0) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 3) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 13) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 2) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 6) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 5) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 7) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 6) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 0) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 8) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 6) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 6) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 4) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 1) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 7) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 14) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 12) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 1) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 6) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 7) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 4) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 1) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 14) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 5) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 3) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 15) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 12) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 7) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 4) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 14) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 2) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 6) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 14) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 3) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 11) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 12) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 0) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 11) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 11) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 14) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 1) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 2) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 11) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 3) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 11) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 0) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 12) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 8) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 11) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 11) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 14) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 15) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 1) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 9) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 5) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 5) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 0) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 14) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 14) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 10) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 6) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 0) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 3) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 14) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 1) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 11) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 14) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 11) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 14) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 1) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 14) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 15) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 12) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 8) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 13) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 11) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 7) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 3) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 7) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 9) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 11) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 8) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 5) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 6) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 8) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 3) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 0) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 12) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 2) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 9) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 0) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 5) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 14) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 7) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 15) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 15) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 1) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 8) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 11) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 7) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 4) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 3) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 15) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 3) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 14) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 1) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 1) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 11) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 0) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 1) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 5) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 12) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 6) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 14) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 8) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 6) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 11) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 5) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 6) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 11) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 12) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 1) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 7) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 5) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 6) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 1) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 7) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 3) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 4) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 13) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 8) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 11) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 1) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 0) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 2) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 9) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 15) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 7) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 12) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 7) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 15) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 3) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 10) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 2) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 11) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 14) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 4) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 9) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 6) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 14) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 12) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 15) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 12) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 14) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 3) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 6) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 6) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 3) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 14) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 9) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 8) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 5) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 8) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 4) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 0) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 15) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 6) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 2) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 8) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 14) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 1) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 10) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 9) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 9) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 5) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 0) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 12) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 14) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 0) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 14) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 3) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 15) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 7) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 14) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 6) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 7) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 6) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 2) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 10) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 0) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 14) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 10) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 2) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 3) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 1) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 6) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 2) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 4) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 0) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 14) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 6) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 6) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 11) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 0) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 11) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 12) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 2) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 6) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 0) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 6) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 12) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 14) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 14) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 13) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 7) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 14) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 3) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 14) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 10) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 7) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 3) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 8) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 15) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 0) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 15) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 5) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 6) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 12) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 14) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 3) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 13) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 12) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 8) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 7) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 6) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 4) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 12) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 15) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 9) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 7) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 5) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 5) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 6) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 15) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 0) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 10) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 6) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 4) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 15) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 8) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 8) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 7) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 1) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 13) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 4) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 7) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 13) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 12) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 5) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 7) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 11) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 12) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 9) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 8) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 9) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 14) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 10) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 1) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 1) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 4) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 12) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 11) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 13) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 3) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 1) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 4) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 8) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 14) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 6) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 7) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 0) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 10) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 14) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 3) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (101); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 4) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 2) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 6) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 11) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 10) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 10) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 11) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 7) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 15) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 12) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (110); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 0) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 14) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 0) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 4) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 15) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 7) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 11) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 1) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 2) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 15) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (112); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 9) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 13) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 1) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 3) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 7) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 2) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (107); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 2) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 6) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 1) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 12)) & progmask

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 13) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (104); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 7) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 2) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 3) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 6) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 5) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 4) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (109); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 14) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 12) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 1) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 3) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 1) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 7) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 5) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 7) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 9) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 5) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 3) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 11) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 10) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 12) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 13) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 12) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 14) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (105); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 6) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = (head + 14) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 12) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 8) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 2) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (105); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 1) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 15) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 5) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 14)) & progmask

		a = 0
		for ; prog[a] != (109); a++ {
		}
		b = 0
		for ; prog[b] != (97); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 6)) & progmask

		a = (head + 9) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (98); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 8) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 1)) & progmask

		a = (head + 4) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 2)) & progmask

		a = (head + 7) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (100); a++ {
		}
		b = 0
		for ; prog[b] != (102); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 0) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 2) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (103); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 13) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 9) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 4) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 10) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 8) & progmask
		b = (head + 3) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 10)) & progmask

		a = (head + 9) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 1) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (97); a++ {
		}
		b = 0
		for ; prog[b] != (107); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 5) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 11) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 3) & progmask
		b = (head + 4) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 8) & progmask
		b = (head + 12) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 3) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (101); a++ {
		}
		b = 0
		for ; prog[b] != (112); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 15) & progmask
		b = (head + 11) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 14) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 12) & progmask
		b = (head + 10) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = (head + 9) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (99); a++ {
		}
		b = 0
		for ; prog[b] != (100); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 11) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 4) & progmask
		b = (head + 7) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (98); a++ {
		}
		b = 0
		for ; prog[b] != (106); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 15)) & progmask

		a = (head + 9) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 11)) & progmask

		a = (head + 0) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 3)) & progmask

		a = (head + 11) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (111); a++ {
		}
		b = 0
		for ; prog[b] != (110); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 14) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (103); a++ {
		}
		b = 0
		for ; prog[b] != (108); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 5)) & progmask

		a = (head + 5) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 10) & progmask
		b = (head + 0) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 8)) & progmask

		a = 0
		for ; prog[a] != (102); a++ {
		}
		b = 0
		for ; prog[b] != (99); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 4) & progmask
		b = (head + 15) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (108); a++ {
		}
		b = 0
		for ; prog[b] != (104); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		a = (head + 8) & progmask
		b = (head + 5) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 4)) & progmask

		a = (head + 0) & progmask
		b = (head + 14) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 9)) & progmask

		a = (head + 8) & progmask
		b = (head + 1) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		a = 0
		for ; prog[a] != (106); a++ {
		}
		b = 0
		for ; prog[b] != (111); b++ {
		}
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 7)) & progmask

		a = (head + 13) & progmask
		b = (head + 6) & progmask
		prog[a], prog[b] = prog[b], prog[a]

		head = (head + (proglen - 13)) & progmask

		a = (head + 0) & progmask
		b = (head + 2) & progmask
		prog[a], prog[b] = prog[b], prog[a]

	}
	end := time.Now()

	fmt.Printf("%dms\n", end.Sub(start)/1000/1000)

	fmt.Printf("Program sequence: ")
	for i := ProgType(0); i < proglen; i++ {
		fmt.Printf("%c", prog[(head+i)%proglen])
	}
	fmt.Printf("\n")
}
