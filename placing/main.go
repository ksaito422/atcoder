package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s1 string
	fmt.Scan(&s1)

	var count int
	arr := strings.Split(s1, "")
	for _, v := range arr {
		i, _ := strconv.Atoi(v)
		if i == 1 {
			count++
		}
	}
	fmt.Println(count)
}
