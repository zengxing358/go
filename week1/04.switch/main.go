package main

import "fmt"

func main() {
	var money int = 20
	var busy bool = true

	switch money {
	case 20:
		fmt.Println("点个外卖")
		fallthrough
	case 100:
		fmt.Println("再考虑")
	case 150:
		fmt.Println("出去看看")
	case 200:
		fmt.Println("下个馆子")
		if busy {
			break
		} else {
			fmt.Println("再吃点零食")
		}

	}
	fmt.Println("end")
}
