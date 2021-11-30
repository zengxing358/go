package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var hello string = "hello, golang!"
	var world = "world"
	fmt.Println(hello, world)
	f := 1.234
	fmt.Println(f)
	var int3, int4 uint = 33, 44 //uint 不带符号的整数
	fmt.Println(int3, int4)

	var h0, ver float64 = 3, 4.123
	var sc = h0 * ver
	fmt.Println(h0 * ver)
	fmt.Println(sc)

	var h1, ver1 float64
	var sc1 = h1 * ver1
	fmt.Println(h1 * ver1)
	fmt.Println(sc1)

	var newname string
	fmt.Println("newname=", newname)

	a1 := "hello"
	fmt.Println(reflect.TypeOf(a1))

	var a6 uint = math.MaxUint64
	var a7 int = int(a6)
	fmt.Println(a6, a7)

	name := "hello"
	age := 30
	fmt.Printf("%p\n", &age)
	age, time := 32, "时间"
	fmt.Printf("%p\n", &age)
	fmt.Println(name, age, time)
	{
		age := 3
		fmt.Printf("%p\n", &age)
	}
}
