package main

// ref: https://go.dev/tour/methods/20

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %.0f", float64(e))
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	precision := 1e-15
	// precision := 0.00005
	half := x / 2
	substract := float64(0)
	old := float64(0)
	add := float64(0)
	loop := 0

	// for: init, evaluate, post
	for sqr := half * half; x-sqr > precision || x-sqr < -1*precision; sqr = half * half {
		// fmt.Printf("%v: %.2f => %.2f \n", x, half, sqr)
		loop += 1
		if sqr > x {
			// fmt.Printf("(%.2f + %.2f) / 2 \n", old, half)
			half = (old + half) / 2
			substract = old
		} else {
			// fmt.Printf("%.2f + ((%.2f - %.2f) / 2)\n", half, half, substract)
			add = ((half - substract) / 2)
			old = half
			half = half + add

		}

		// if loop > 10 {
		// 	break
		// }

	}

	// fmt.Printf("Sqrt: loop %v => %v\n", x, loop)

	return half, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
