package main

import "fmt"

func main() {
	fmt.Println("中午呢")
	fmt.Printf("My name is %s\n", "zx") //直接输出字符串
	fmt.Printf("My name is %q\n", "zx") //双引号括起来的字符串
	fmt.Printf("My name is %x\n", "zx") //每个字节用两字符16进制 （a-f)
	fmt.Printf("My name is %X\n", "zx") //每个字节用两字符16进制 （A-F)

	fmt.Printf("%.3f", 1.23455676)
}
