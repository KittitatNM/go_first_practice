package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Saturday

	switch today {
	case time.Monday:
		fmt.Println("Today is Monday")
	case time.Saturday, time.Sunday:
		fmt.Println("Today is Weekend")
		fallthrough
	default:
		fmt.Println("Today is fine")
	}
}
