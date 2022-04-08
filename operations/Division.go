package operations

import (
	"fmt"
	"strconv"
)

func Div(c int64) int64 {
	fmt.Println("Enter the number")
	var num string
	fmt.Scan(&num)
	if num == "esc" {
		return 0
	}
	test, err := strconv.ParseInt(num, 0, 64)
	if err != nil {
		fmt.Println("string conversion error")
	}
	fmt.Println(num)
	c = c / test
	test = 0
	return c
}
