package main

import "fmt"

func main() {
	//1D array
	var arr [3][3]int
	var x int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Enter the %d%d position element:", i+1, j+1)
			fmt.Scan(&x)
			arr[i][j] = x
		}
	}
	fmt.Println(arr)
}
