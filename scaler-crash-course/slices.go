package main

import "fmt"

func slices(){

	// create a slice that stores 3 elements
	s := make([]string, 3)
	fmt.Println(s)

	// store elements in the slice
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s)

	// append value 'd' to the slice
	s = append(s, "d")
	fmt.Println(s)
	fmt.Println(len(s))

	// iterate over & access silces

	for i, num := range(s) {
		fmt.Println(i,num)
	}
}



func main(){
	slices()
}