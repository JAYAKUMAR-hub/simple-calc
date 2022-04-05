package operations

import (
	"fmt"
	"strconv"
)

var num string

func Sub(c int64) int64 {
	fmt.Print("Enter number: ")
	fmt.Scan(&num)
	if num == "esc" {
		return 0
	}
	test, err := strconv.ParseInt(num, 0, 64)
	if err != nil {
		fmt.Println("str conversion error")
	}
	c = c - test
	test = 0
	return c
}
