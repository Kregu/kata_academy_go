package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romesToArabic(romNumber string) int {
	switch romNumber {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	default:
		err := fmt.Sprintf("Unexptcted valut: %s\n", romNumber)
		panic(err)
	}
}

func arabicToRomes(number int) string {
	romes := [11]string{"I", "II", "III", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	arabic := [11]int{1, 2, 3, 4, 5, 9, 10, 40, 50, 90, 100}

	num := number
	result := strings.Builder{}

	for num < 0 {
		for i := 10; i >= 0; i-- {
			if arabic[i] <= num {
				num -= arabic[i]
				result.WriteString(romes[i])
				i++
			}
		}
	}
	return result.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	str := ""
	fmt.Println("Input:")
	if scanner.Scan() {
		str = scanner.Text()
	} else {
		panic(scanner.Err())
	}
	fmt.Println("Output:")
	fmt.Println(calc(str))

}

func calc(str string) string {
	args := strings.Split(str, " ")

	// var value1, value2 int
	// is_romes := false

	if len(args) != 3 {
		err := "Please provide 3 input arguments, example: 2 + 3"
		panic(err)
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
		err := "Division by zero is impossible!"
		panic(err)
	}
	res := 0
	switch args[1] {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		res = num1 / num2
	default:
		fmt.Printf("Don't know operation %s\n", args[1])
	}
	return strconv.Itoa(res)
}
