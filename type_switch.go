// package main

// import "fmt"

// func main() {
// 	whoAmI := func(i interface{}) {
// 		switch t := i.(type) {
// 		case int:
// 			fmt.Println("this is int", t)
// 		case bool:
// 			fmt.Println("this is boolean", t)
// 		case string:
// 			fmt.Println("this is string", t)

// 		}
// 	}
// 	//whoAmI(10)
// 	whoAmI("Hello")
// }