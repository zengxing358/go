package main

import "fmt"

func main() {
	var rmb int = 21
	if rmb <= 20 {
		fmt.Println("点个外卖")
	} else if rmb <= 200 {
		fmt.Println("下个馆子")
	} else if rmb <= 2000 {
		fmt.Println("去米其林")
	} else if rmb <= 20000 {
		fmt.Println("去新马泰")
	} else {
		fmt.Println("容我想想")
	}
	fmt.Println("end")
}
