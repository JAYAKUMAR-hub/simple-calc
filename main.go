package main

import (
	"calculator/operations"
	"fmt"
	"strconv"
)

var c int64 = 0
var d int64 = 0
var b string
var a string
var flag int = 0

func Num() int64 {
	for {
		fmt.Print("Enter number: ")
		fmt.Scan(&a)
		if a == "esc" {
			d = 0
			continue
		}
		test, err := strconv.ParseInt(a, 0, 64)
		if err != nil {
			fmt.Println("str conversion error")
			continue
		}
		if c == 0 {
			c = c + test
			test = 0
			fmt.Println(c)
		} else {
			Symbol()
		}
		return c
	}
	return 0
}

func Eval(b string) int64 {
	if b != "=" {
		switch b {
		case "+":
			d += operations.Add(c)
		case "-":
			fmt.Println(d)

			d += operations.Sub(c)
			if d == 0 {
				return 0
			}
		case "*":
			if d == 0 {
				d = operations.Mul(c)
			} else {
				d = operations.Mul(d)
			}
		case "/":
			if d == 0 {
				d = operations.Div(c)
			} else {
				d = operations.Div(d)
			}
		case "esc":
			flag = 1
			d = 0
			break
		default:
			fmt.Println("Invalid")
		}
	} else if b == "=" {
		fmt.Println("The result is: ", d)
	}
	return d
}

func Symbol() int64 {
	for {
		fmt.Println("Enter the symbol")
		fmt.Scan(&b)
		d = Eval(b)
		if d == 0 && flag == 1 {
			d = 0
			break
		}
		c = 0
	}
	return d
}

func Submain() {
	for {
		Num()
		Symbol()
	}
}

func main() {
	Submain()
}
