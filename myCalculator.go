package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var expectedOperations = []string{"+", "-", "*", "/"}

var romanNumeralMap = map[string]int{
	"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
}

func romanToArabic(roman string) int {

	arabic := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumeralMap[string(roman[i])]
		if value < prevValue {
			arabic -= value
		} else {
			arabic += value
		}
		prevValue = value
	}
	return arabic
}

func arabicToRoman(arab int) string {

	arabic := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; i < len(arabic); i++ {
		for arab >= arabic[i] {
			arab -= arabic[i]
			result += romans[i]
		}
	}
	return result
}

func contains(mass []string, str string) bool {
	for _, value := range mass {
		if value == str {
			return true
		}
	}
	return false
}

func isRoman(str string) bool {
	for num := 0; num < len(str); num++ {
		letter := string(str[num])
		_, exists := romanNumeralMap[letter]
		if !exists {
			return false
		}
	}
	return true
}

func validate(left string, oper string, right string) bool {

	_, left_err := strconv.Atoi(left)
	_, right_err := strconv.Atoi(right)
	isOper := contains(expectedOperations[:], oper)

	if left_err == nil && right_err == nil && isOper {
		if oper == expectedOperations[3] && right == "0" {
			return false
		}
		return true
	} else if isRoman(left) && isRoman(right) && isOper {
		return true
	}
	return false
}

func sum(x, y int) int {
	sum := x + y
	return sum
}

func sub(x, y int) int {
	sub := x - y
	return sub
}

func mul(x, y int) int {
	mul := x * y
	return mul
}

func div(x, y int) int {
	div := x / y
	return div
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.Trim(line, "\n")

	parts := strings.Split(line, " ")

	if len(parts) == 3 && validate(parts[0], parts[1], parts[2]) {
		oper1 := 0
		oper2 := 0
		if isRoman(parts[0]) {
			oper1 = romanToArabic(parts[0])
			oper2 = romanToArabic(parts[2])
		} else {
			oper1, _ = strconv.Atoi(parts[0])
			oper2, _ = strconv.Atoi(parts[2])
		}

		if oper1 <= 10 && oper2 <= 10 {

			operation := parts[1]
			result := 0

			if operation == expectedOperations[0] {
				result = (sum(oper1, oper2))
			} else if operation == expectedOperations[1] {
				result = (sub(oper1, oper2))
			} else if operation == expectedOperations[2] {
				result = (mul(oper1, oper2))
			} else if operation == expectedOperations[3] {
				result = (div(oper1, oper2))
			}
			if isRoman(parts[0]) {
				if result < 0 {
					fmt.Println("В римской системе нет отрицательных чисел")
				} else if result == 0 {
					fmt.Println("В римской системе нет 0")
				} else {
					fmt.Println(arabicToRoman(result))
				}
			} else {
				fmt.Println(result)
			}
		} else {
			fmt.Println("Введены числа больше 10")
		}
	} else {
		fmt.Println("Введенны некорректные данные")
	}
}
