package main

import "fmt"

func arr_way1(){
	var a[5] int 
	fmt.Println(a)

	var b[] int 
	fmt.Println(b)

	a[3] = 100
	fmt.Println(a)
}

func arr_way2(){
	a := [5] int {1,2,3,4,5}
	fmt.Println(a)
}

func main(){
	// arr_way1()
	arr_way2()
}