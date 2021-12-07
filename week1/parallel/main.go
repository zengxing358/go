package main

import (
	"fmt"
)

func main() {
	lines := [][][]float64{}
	types := []string{"start", "end"}
	for i := 1; i <= 2; i++ {
		line := [][]float64{}
		for j := 1; j <= 2; j++ {
			var x float64
			fmt.Printf("line%d %s点x：", i, types[0])
			fmt.Scanln(&x)

			var y float64
			fmt.Printf("line%d %s点y：", i, types[1])
			fmt.Scanln(&y)
			line = append(line, []float64{x, y})
		}
		lines = append(lines, line)
	}
	if checkParalleLine(lines) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

func checkParalleLine(lines [][][]float64) bool {
	x0 := lines[0][0][0]
	y0 := lines[0][0][1]
	x1 := lines[0][1][0] - x0
	y1 := lines[0][1][1] - y0
	x2 := lines[1][0][0]
	y2 := lines[1][0][1]
	x3 := lines[1][1][0] - x2
	y3 := lines[1][1][1] - y2
	if x1*y3 != x3*y1 {
		return false
	}
	return true
}
