package main

import "fmt"

func main() {
	var age int
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)
	if age < 18 {
		fmt.Println("you can not drive")
	} else {
		fmt.Println("you can drive")
	}
	fmt.Printf("your age was %d", age)
}
