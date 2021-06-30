package main

import "fmt"

func main() {
	fmt.Println("hello,world")
	case1()
	case2()
	case3()
	case4()
	case5()
}
func case1() {
	var a = 1
	fmt.Println(a)
}
func case2() {
	a := 1
	fmt.Println(a)
}
func case3() {
	a, b, c := 1, 2, 3
	fmt.Println(a, b, c)
}
func case4() {
	var (
		a = 0
		b = "0"
		c = 0.1
	)
	fmt.Println(a, b, c)
}

func case5() {
	const (
		a = iota << 10
		b
		c
		d
	)
	fmt.Println(a, b, c, d)
}
