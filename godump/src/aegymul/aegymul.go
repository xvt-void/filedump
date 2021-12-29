package main

import (
	"flag"
	"fmt"
)

func main() {

	p := flag.Uint64("p", 3, "")
	q := flag.Uint64("q", 4, "")

	flag.Parse()

	fmt.Println(aegymul(*p, *q))

}

func aegymul(p uint64, q uint64) uint64 {

	var e uint64 = 0

	if p < q {

		for p >= 1 {
			if p%2 == 1 {
				e = e + q
			}

			p = p / 2
			q = q * 2

		}
	} else {

		for q >= 1 {
			if q%2 == 1 {
				e = e + p
			}

			q = q / 2
			p = p * 2

		}
	}

	return e
}
