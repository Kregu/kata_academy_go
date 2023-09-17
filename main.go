package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Print("Please provide 3 input arguments, example: 2 + 3")
		return
	}

	num1, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}
	num2, err := strconv.Atoi(args[2])
	if err != nil {
		panic(err)
	}
	if num2 == 0 && (args[1] == "/" || args[1] == "%") {
		fmt.Print("Division by zero is impossible!")
		return
	}

	switch args[1] {
	case "+":
		fmt.Printf("result: %d\n", num1+num2)
	case "-":
		fmt.Printf("result: %d\n", num1-num2)
	case "*":
		fmt.Printf("result: %d\n", num1*num2)
	case "/":
		fmt.Printf("result: %.6f\n", float64(num1)/float64(num2))
	case "%":
		fmt.Printf("result: %d\n", num1%num2)
	default:
		fmt.Printf("Don't know operation %s\n", args[1])
	}
}
