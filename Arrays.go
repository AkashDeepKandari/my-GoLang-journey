package main

import "fmt"

func main() {
	//1D array
	var arr [4]int
	var x int
	for i := 0; i < 4; i++ {
		fmt.Printf("Enter the %d element:", i+1)
		fmt.Scan(&x)
		arr[i] = x
	}

	fmt.Println(arr)
}
