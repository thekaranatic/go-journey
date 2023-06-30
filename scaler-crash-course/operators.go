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


func main(){
	incr_decr()
	conjunction_op()
}