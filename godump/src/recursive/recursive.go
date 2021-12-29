package main

import (
	"flag"
	"fmt"
)

func main() {

	modeptr := flag.String("mode", "", "add|pot|mul")
	mptr := flag.Uint64("m", 2, "value")
	nptr := flag.Uint64("n", 8, "recursion")

	flag.Parse()

	mode := *modeptr
	m := *mptr
	n := *nptr

	if mode == "add" {

		fmt.Printf("[1 + add(m, n -1)]\n------------------\n")
		result := add(m, n)
		fmt.Printf("------------------\n%d + %d = %d", m, n, result)
	} else if mode == "pot" {
		fmt.Printf("[m * pot(m, n -1)]\n------------------\n")
		result := pot(m, n)
		fmt.Printf("------------------\n%d ^ %d = %d", m, n, result)
	} else if mode == "mul" {
		fmt.Printf("[m + mul(m, n -1)]\n------------------\n")
		result := mul(m, n)
		fmt.Printf("------------------\n%d * %d = %d", m, n, result)
	} else {
		fmt.Println("use '-h' for help")
		return
	}
}

func add(m uint64, n uint64) uint64 {

	if n == 0 {
		return (m)
	} else {
		fmt.Printf("[1 + add(%d, %d -1)]\n", m, n)
		r_ := 1 + (add(m, n-1))
		fmt.Printf("[1 + add(%d)] = %d\n", m+(n-1), r_)
		return (r_)
	}

}

func pot(m uint64, n uint64) uint64 {

	if n == 1 {
		return (m)
	} else {
		fmt.Printf("[%d * pot(%d, %d -1)]\n", m, m, n)
		r_ := m * (pot(m, n-1))
		fmt.Printf("[%d * pot(%d)] = %d\n", m, n, r_)
		return (r_)
	}
}

func mul(m uint64, n uint64) uint64 {
	if n == 1 {
		return (m)
	} else {
		fmt.Printf("[%d + mul(%d, %d -1)]\n", m, m, n)
		r_ := m + (mul(m, n-1))
		fmt.Printf("[%d + mul(%d * %d)] = %d\n", m, n-1, m, r_)
		return (r_)
	}
}

// 2*8 results in 7row & 7row bc 0 is not included, there may be a workaround
// missing -> [2 + mul(0*2)] = 2
// possible solution m < n
