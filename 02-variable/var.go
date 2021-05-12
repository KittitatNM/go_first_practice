package main

import "fmt"

func main() {
	// var name string = "Gopher นำแน่"
	var name = "Gopher นำแน่"
	// name := "Gopher นำแน่"

	fmt.Printf("Name is %v\n", name)

	fmt.Printf("type is : %T\n", name) // แสดง Typeof ของตัวแปร
	fmt.Printf("type is : %q\n", name) // แสดง qoute ของข้อมูล

	// constant
	const day = "Monday"
	fmt.Printf("Today is %v\n", day)

	// const (
	// 	Sunday = 0
	// 	Monday = 1
	// )

	const (
		Sunday = iota
		Monday
	)

	fmt.Println("Sunday is ", Sunday)
	fmt.Println("Monday is ", Monday)

}
