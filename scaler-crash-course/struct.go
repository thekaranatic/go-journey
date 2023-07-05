package main

import "fmt"

type closefriends struct{
	name string
	domain string
	age int
}

func close_friends(){
	// enter close friends of Karan in A-Z order
	fmt.Println(closefriends{"abc","data science", 21})
	fmt.Println(closefriends{"def","computer science", 20})
	fmt.Println(closefriends{"ghi","graphic design", 21})
	fmt.Println(closefriends{"jkl","dental science", 20})
	fmt.Println(closefriends{"mno","computer science", 20})
}

type rect struct{
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}


func main(){

	r := rect{width:17, height:3}
	fmt.Println("Area:", r.area())
}