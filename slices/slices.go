package main

import "tour-of-go/slices/pic"

func Pic(dx, dy int) [][]uint8 {
	height := dy
	width := dx
	PIC := make([][]uint8, height) // create a slice of slices

	for i := 0; i < height; i++ {
		picRow := make([]uint8, width) // create a slice for each row
		for j := 0; j < width; j++ {
			picRow[j] = uint8(j*width ^ height*i)
		}
		PIC[i] = picRow
	}
	return PIC
}
func main() {
	pic.Show(Pic)
}
