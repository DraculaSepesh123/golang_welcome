package main

import (
	"fmt"
	"strconv"
)

func main() {

	var roman = [100]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}

	var workingWithRoman bool

	var number1, operation, number2 string
	_, err := fmt.Scanln(&number1, &operation, &number2)
	if err != nil {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}

	if number1 == "" || number2 == "" || operation == "" {
		fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
		return
	}

	convnum1, err_1 := strconv.Atoi(number1)
	convnum2, err_2 := strconv.Atoi(number2)

	if err_1 != nil && err_2 != nil {
		workingWithRoman = true
		for z := 0; z < 100; z += 1 {
			if number1 == roman[z] {
				convnum1 = z + 1
				break
			}
		}

		for z := 0; z < 100; z += 1 {
			if number2 == roman[z] {
				convnum2 = z + 1
				break
			}
		}
	} else if err_1 == nil && err_2 == nil {
		workingWithRoman = false
	} else {
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		return
	}

	var res int

	if operation == "+" {
		res = convnum1 + convnum2
	} else if operation == "-" {
		res = convnum1 - convnum2
	} else if operation == "*" {
		res = convnum1 * convnum2
	} else if operation == "/" {
		res = convnum1 / convnum2
	} else {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}

	if workingWithRoman {
		if res <= 0 {
			fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
			return
		}
		fmt.Println(roman[res-1])
	} else {
		fmt.Println(res)
	}
}
