package operations

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Calcuator() {
	fmt.Print("Enter number: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	result, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println(result)

	for {
		var symbol string
		var number int64
		symbol = scanner.Text()
		scanner.Scan()
		if symbol == "=" {
			break
		} else if symbol != "+" || symbol != "-" || symbol == "*" || symbol == "/" {
			fmt.Println("invalid")
		}
		scanner.Scan()
		number, _ = strconv.ParseInt(scanner.Text(), 10, 64)

		result = calculation(result, number, symbol)

	}
	fmt.Println(result)
}

func calculation(result int64, number int64, symbol string) int64 {
	if symbol == "+" {
		return result + number
	} else if symbol == "-" {
		return result - number
	} else if symbol == "*" {
		return result * number
	} else if symbol == "/" {
		return result / number
	}
	return result
}
