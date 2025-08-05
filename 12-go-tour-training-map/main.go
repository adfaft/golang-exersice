package main

// ref: https://go.dev/tour/moretypes/23

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	result := make(map[string]int)

	fields := strings.Fields(s)
	for _, i := range fields {
		_, ok := result[i]
		if ok {
			result[i] += 1
		} else {
			result[i] = 1
		}
	}

	return result
}

// ref: https://gist.github.com/zyxar/2317744
func WordCountReal(s string) map[string]int {
	ss := strings.Fields(s)
	num := len(ss)
	ret := make(map[string]int)
	for i := 0; i < num; i++ {
		(ret[ss[i]])++
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
