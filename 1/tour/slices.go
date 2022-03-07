package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	output := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			output[y] = append(output[y], uint8(x^y))
		}
	}

	return output
}

func main() {
	pic.Show(Pic)
}
