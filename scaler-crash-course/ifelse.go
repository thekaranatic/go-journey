package main

import "fmt"

func ifElse(){
	var a, b int = 15, 25

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
}

func main(){
	ifElse()
}