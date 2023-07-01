package main

import "fmt"

func ifElse(){
	sum := 10

	if sum < 10 {
		fmt.Println("less than 10")
	} else {
		fmt.Println("greater than 10")
	}
}


func ifElse_one(){
	sum := 10
	sum := 20

	if sum == 40 {
		fmt.Println("equal to 40")
	} else {
		fmt.Println("soem other value other than 40")
	}
}

func main(){
	ifElse()
	ifElse_one()
}