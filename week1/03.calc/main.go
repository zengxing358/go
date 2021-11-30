package main

import "fmt"

func main() {
	var a, b int = 3, 10
	fmt.Println("a+b=", a+b)
	fmt.Println("a-b=", a-b)
	fmt.Println("a*b=", a*b)
	fmt.Println("a/b=", a/b)
	fmt.Println("a%b=", a%b)

	fmt.Println(true && false)
	fmt.Println(true || false)

	fmt.Println(a > b)
	fmt.Println(a < b)

}
