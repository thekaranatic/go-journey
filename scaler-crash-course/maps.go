package main

import "fmt"

func maps(){

	// creates a map with 
	// strings as type of Keys and int that of values
	m := make(map[string] int)

	m["age"] = 21
	m["dob"] = 28

	fmt.Println(m)

	delete(m,"age")
	fmt.Println(m)
}

func maps_one(){

	// creates a map with 
	// strings as type of Keys and int that of values
	m := map[string] int {"age":21, "dob":28}

	fmt.Println(m)

	delete(m,"age")
	fmt.Println(m)
}

func main(){
	maps()
	maps_one()
}