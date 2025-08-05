package main

// erf: https://go.dev/tour/moretypes/18

import "golang.org/x/tour/pic"

// import "fmt"

func Pic(dx, dy int) [][]uint8 {

	var result [][]uint8

	var temp []uint8
	for i := range dy {
		temp := temp[:0]
		for j := range dx {
			temp = append(temp, uint8(i*j))
		}
		result = append(result, temp)
	}

	return result

}

// ref : https://gist.github.com/zyxar/2317744
func PicReal(dx, dy int) [][]uint8 {
	var result = make([][]uint8, dy)
	for x := range result {
		result[x] = make([]uint8, dx)
		for y := range result[x] {
			result[x][y] = uint8(x * y)
		}
	}
	return result
}

func main() {
	pic.Show(Pic)
	//fmt.Println(Pic(7,5))
	//fmt.Println(PicReal(7,5))
}
