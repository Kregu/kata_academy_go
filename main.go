package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romesToArabic(romNumber string) int {
	switch strings.ToUpper(romNumber) {
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

	result := strings.Builder{}

	for number > 0 {
		for i := 10; i >= 0; i-- {
			if arabic[i] <= number {
				number -= arabic[i]
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

	isNum1romes := false
	isNum2romes := false

	if len(args) != 3 {
		err := "Requires two operands and one operator (+, -, /, *). Example: 1 + 2"
		panic(err)
	}

	num1, err1 := strconv.Atoi(args[0])
	if err1 != nil {
		num1 = romesToArabic(args[0])
		isNum1romes = true
	}

	num2, err2 := strconv.Atoi(args[2])
	if err2 != nil {
		num2 = romesToArabic(args[2])
		isNum2romes = true
	}

	if isNum1romes != isNum2romes {
		panic("Please use only arabic or only roman numbers")
	}

	if num1 < 1 || num2 < 1 || num1 > 10 || num2 > 10 {
		panic("Arguments must be: 1..10")
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
		err := fmt.Sprintf("Don't know operation %s\n", args[1])
		panic(err)
	}

	if isNum1romes {
		if res > 0 {
			return arabicToRomes(res)
		} else {
			err := fmt.Sprintf("Roman numerals cannot be less than one: %d", res)
			panic(err)
		}
	}

	return strconv.Itoa(res)
}
