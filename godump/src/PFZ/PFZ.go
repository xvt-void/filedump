package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {

	m := flag.Uint64("m", 988027, "Produkt aus 2 Primzahlen -> 991*997 = 988027")

	flag.Parse()

	fmt.Println("m:", *m)

	if (*m % 2) == 0 {

		fmt.Printf("%d", 2)

	} else if (*m % 3) == 0 {

		fmt.Printf("%d", 3)
	}

	var z float64 = 5
	var g float64 = math.Sqrt(float64(*m)) //does not have to be a float but needs to be truncated

	for z < g {
		if ((uint64(*m)) % (uint64(z))) == 0 {
			var p uint64 = uint64(z)
			var q uint64 = (uint64(*m) / p)

			//fmt.Printf("%d\n", uint64(z))
			fmt.Printf("p: %d\n", p)
			fmt.Printf("q: %d\n", q)

		}

		z += 2
	}

}
