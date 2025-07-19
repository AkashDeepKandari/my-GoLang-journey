package main

import (
	"fmt"
	"slices"
)

func main() {
	//A Slice is a dynamic sized, flexible view into an array
	//make()->Used to initialize slices, maps, or channels with memory allocation.
	//var sl []int {2,3,5,...}
	var sl = make([]int, 0)
	sl = append(sl, 1)
	sl = append(sl, 3)
	sl = append(sl, 6)
	sl = append(sl, 8)
	sl = append(sl, 23)
	fmt.Println(sl)
	//copy method
	var sl1 = make([]int, len(sl))
	copy(sl1, sl)
	fmt.Println(sl1)
	//cap method
	fmt.Println(cap(sl))
	//equal func
	fmt.Println(slices.Equal(sl, sl1)) //bool value
}
