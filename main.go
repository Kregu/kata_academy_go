package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romesToArabic(romNumber string) (int, error) {
	switch strings.ToUpper(romNumber) {
	case "I":
		return 1, nil
	case "II":
		return 2, nil
	case "III":
		return 3, nil
	case "IV":
		return 4, nil
	case "V":
		return 5, nil
	case "VI":
		return 6, nil
	case "VII":
		return 7, nil
	case "VIII":
		return 8, nil
	case "IX":
		return 9, nil
	case "X":
		return 10, nil
	default:
		return 0, fmt.Errorf("unexptcted value: %s", romNumber)
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

func calc(str string) (string, error) {
	args := strings.Split(str, " ")

	if len(args) != 3 {
		return "", fmt.Errorf("requires two operands and one operator (+, -, /, *). Exam: 1 + 2")
	}

	num1, err1 := strconv.Atoi(args[0])
	if err1 != nil {
		num1, err1 = romesToArabic(args[0])
		if err1 != nil {
			return "", err1
		}
	}

	num2, err2 := strconv.Atoi(args[2])
	if err2 != nil {
		num2, err2 = romesToArabic(args[2])
		if err2 != nil {
			return "", err2
		}
	}

	if num1 < 1 || num2 < 1 || num1 > 10 || num2 > 10 {
		return "", fmt.Errorf("arguments must be: 1..10")
	}

	if num2 == 0 && (args[1] == "/" || args[1] == "%") {
		return "", fmt.Errorf("division by zero is impossible!")
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
		return "", fmt.Errorf("don't know operation %s", args[1])
	}

	if res > 0 {
		return arabicToRomes(res), nil
	}

	return strconv.Itoa(res), nil
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
	res, err := calc(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
