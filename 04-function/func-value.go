// package main

// import "fmt"

// var plus func(a, b int) int = func(a, b int) int { return a + b }

// func cal(op func(int, int) int) {
// 	result := op(4, 5)
// 	fmt.Println("Cal result is : ", result)
// }

// func main() {
// 	p := plus(1, 2)
// 	fmt.Println(p)
// 	fmt.Printf("Type of %T\n", plus)
// 	cal(plus)

// 	minus := func(a, b int) int { return a - b }
// 	cal(minus)
// }
