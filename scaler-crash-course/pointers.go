package main

import "fmt"

func main() {
	i := 28
	
	// REFERENCING: point to the address of `i` and store the same into`p`
	p := &i 	

	fmt.Println(i)
	fmt.Println(p)	// prints memory location/address of i

	// DEREFERENCING: print the value in the memory location pointed at
	fmt.Println(*p)

	// change the value in the memory location pointed at
	*p = 21
	fmt.Println(i)

}