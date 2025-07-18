// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("What is your name?")
// 	x, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("Error reading input:", err)
// 		return
// 	}
// 	name := strings.TrimSpace(x)
// 	fmt.Print(name)
// }