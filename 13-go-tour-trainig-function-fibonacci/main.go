package main

// ref: https://go.dev/tour/moretypes/26

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	last, new := 0, 1

	return func() int {
		temp := new + last
		old := last
		last = new
		new = temp
		return old
	}
}

// ref: https://medium.com/@anumsarmadmalik/11-solutions-togolang-tours-exercises-7ee61b7b94f5
func fibonacciReal() func() int {
	total, nextTotal := 0, 1
	return func() int {
		result := total
		total, nextTotal = nextTotal, nextTotal+result
		return result
	}
}

// ref: https://gist.github.com/zyxar/2317744
func fibonacciReal2() func() int {
	var x, y = 0, 1
	return func() (z int) {
		z, x, y = x, y, x+y
		return
	}
}

func main() {
	f := fibonacci()
	// f := fibonacciReal()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
