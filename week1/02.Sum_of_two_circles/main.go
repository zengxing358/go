package main

import "fmt"

func main() {
	var a, b float64
	const pi float64 = 3.1415926
	fmt.Println("请输入2个圆的半径")
	fmt.Scanf("%f %f", &a, &b)
	fmt.Scanf("%f", &b)
	fmt.Printf("两个圆的面积之和是%.3f\n", pi*a*a+pi*b*b)
}
