package main

import (
	"fmt"
	"reflect"
)

func decrl(){
	var i int = 20 // declration & initialisation
	var s string = "golang"
	fmt.Println(i, s)
	fmt.Println()
}

func declr_2(){

	// Go inference - understands types of vars without
	// explicit declaration
	var i = 20 
	var s = "golang"

	fmt.Println(i, reflect.TypeOf(i))
	fmt.Println(s, reflect.TypeOf(s))
	fmt.Println()

}

func chained_assmnt(){
	var fname, lname string = "karan", "kakati"
	var age, weight int = 21, 65

	fmt.Println(fname + lname)
	fmt.Println(age, weight)
	fmt.Println()

}

func walrus(){
	// walrus operator
	name := "karan"
	fmt.Println(name)
	fmt.Println()

}

func float_32(){
	var a,b float32 = 4.5, 5.5
	fmt.Println(a+b)
	fmt.Println(b-a)
	fmt.Println(b/a)
	fmt.Println(b*a)
}


func main(){
	// decl()
	// decl_2()
	// chained_assmnt()
	// walrus()
	float_32()
}