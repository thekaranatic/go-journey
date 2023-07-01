package main

import "fmt"

func incr_decr(){
	var a, b int = 20, 5

	fmt.Println(a%b)
	a++
	fmt.Println(a)

	a -= 1
	fmt.Println(a)
	
}

func conjunction_op(){

	var a int = 20

	a += 1
	fmt.Println(a)
	a /= 1
	fmt.Println(a)
	a *= 1
	fmt.Println(a)
	a %= 1
	fmt.Println(a)
}

func comp_oprs(){
	var a, b int = 15, 25

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
}

func logical_oprs(){
	var x, y, z = 10, 20, 30
	fmt.Println(x<y && z<y)
	fmt.Println(x<y || z>y)
	fmt.Println(!(x<y || z>y))
}

func main(){
	// incr_decr()
	// conjunction_op()
	// comp_oprs()
	logical_oprs()
}