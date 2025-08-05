package main

// ref: https://go.dev/tour/flowcontrol/8

import (
	"fmt"
	"math"
)

/**
0.5  =>
0.75 => 0.5 + ( 0.5  - 0 / 2)
1.12 => 0.75 + ( 0.75 - 0  / 2)
0.94 => (1.12 + 0.75) / 2
0.xx => 0.94 + ( 1.12 - 0.94 ) / 2
*/

func Sqrt(x float64) float64 {
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

	fmt.Printf("Sqrt: loop %v => %v\n", x, loop)

	return half
}

// ref: https://medium.com/@anumsarmadmalik/11-solutions-togolang-tours-exercises-7ee61b7b94f5
func SqrtReal(x float64) float64 {
	z := float64(1)
	for i := 1; i <= 10; i++ {
		// fmt.Println(z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func SqrtReal2(x float64) float64 {
	z := float64(2.)
	s := float64(0)
	loop := 0
	for {
		z = z - (z*z-x)/(2*z)
		loop += 1
		if math.Abs(s-z) < 1e-15 {
			break
		}
		s = z
	}

	fmt.Printf("SqrtReal2: loop for %v => %v\n", x, loop)

	return s
}

func main() {
	fmt.Println(Sqrt(1))    // 1
	fmt.Println(Sqrt(4))    // 2
	fmt.Println(Sqrt(9))    // 3
	fmt.Println(Sqrt(1000)) // 31.62
	fmt.Println(Sqrt(1024)) // 32

	fmt.Println("----")

	fmt.Println(SqrtReal(5))  // 2.24
	fmt.Println(Sqrt(5))      // 2.24
	fmt.Println(SqrtReal2(5)) // 2.24
}
