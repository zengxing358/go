package main

import "fmt"

func main() {
	coordinates := [][]float64{}
	for {
		var x float64
		fmt.Print("点x：")
		fmt.Scanln(&x)

		var y float64
		fmt.Print("点y：")
		fmt.Scanln(&y)
		coordinates = append(coordinates, []float64{x, y})
		var wetherContinue string
		fmt.Print("是否录入下一个（y/n）？")
		fmt.Scanln(&wetherContinue)
		if wetherContinue != "y" {
			break
		}
	}
	if checkStraightLine(coordinates) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

func checkStraightLine(coordinates [][]float64) bool {
	x0 := coordinates[0][0]
	y0 := coordinates[0][1]
	x1 := coordinates[1][0] - x0
	y1 := coordinates[1][1] - y0
	for i := 2; i < len(coordinates); i++ {
		y := coordinates[i][1] - y0
		x := coordinates[i][0] - x0
		if x1*y != x*y1 {
			return false
		}
	}
	return true
}
