package main

import (
	"fmt"
)

func main() {
	var a, b uint
	fmt.Scanf("%d %d", &a, &b)
	res := a * b

	if res%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
