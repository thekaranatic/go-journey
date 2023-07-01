// there is no while loop in Go

package main

import "fmt"

func forLoop_one(){
	sum := 0;

	for i := 0; i<10; i++ {
		sum++;
	}

	fmt.Println(sum)
}

// In Go, you can skip the iterator and initialiser as well in
func forLoop_two(){
	sum := 1

	for sum < 10 {
		sum++
	}

	fmt.Println(sum)
}

func foreverLoop(){
	// infinite loop
	sum := 0

	for{
		sum++
	}

	fmt.Println(sum)
}

func main(){
	// forLoop_one()
	// forLoop_two()
	foreverLoop()
}