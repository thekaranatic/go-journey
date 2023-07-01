package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sayHello(){
	// print "world" before "Hello"
	defer fmt.Println("world")	
	fmt.Println("Hello")
}

func main(){
	sayHello()
	fmt.Println(add(50,50))
}