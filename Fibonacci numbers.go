package main

import (
	"fmt"
	"math"
)

func main() {
	const a = 1.61803398874989
	const b = -1 / a
	var m float64

	for n := 0.0; n < 12.0; n++ {

		m = (math.Pow(a, n) - math.Pow(b, n)) / math.Sqrt(5)
		m = math.Round(m)
		fmt.Println(m)
	}
}
