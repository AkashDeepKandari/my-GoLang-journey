package main

import "fmt"

func main() {
	//A Slice is a dynamic sized, flexible view into an array
	//make()->Used to initialize slices, maps, or channels with memory allocation.
	//var sl []int {2,3,5,...}
	var sl = make([]int, 3)
	sl = append(sl, 1)
	sl = append(sl, 3)
	sl = append(sl, 6)
	fmt.Println(sl)

}
