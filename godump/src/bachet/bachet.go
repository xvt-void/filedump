package main

import (
	"flag"
	"fmt"
)

func main() {

	d := flag.Uint64("d", 7, "")
	m := flag.Uint64("m", 30, "")

	flag.Parse()

	// 1 = x*m+y*d

	bachet(*d, *m)

}

func bachet(d uint64, m uint64) uint64 {

	var x uint64
	var y uint64

	x = m / d
	y = m % d

	for (y != 1) && (y != 0) {

		fmt.Printf("%d = %d * %d + %d\t", m, x, d, y)
		fmt.Printf("\t%d = %d - %d * %d\n", y, m, x, d)

		m = d
		d = y
		x = m / d
		y = m % d

		if y == 1 {
			fmt.Printf("%d = %d * %d + %d\t", m, x, d, y)
			fmt.Printf("\t%d = %d - %d * %d\n", y, m, x, d)
		}
	}

	return 0
}
