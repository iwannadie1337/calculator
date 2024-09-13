package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Карта римских цифр и их значений
var romanToIntMap = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var intToRomanMap = []string{
	"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
}

func isRoman(s string) bool {
	_, exists := romanToIntMap[s]
	return exists
}

func romanToInt(s string) int {
	return romanToIntMap[s]
}

func intToRoman(num int) string {
	if num <= 0 || num > 3999 {
		panic("Невозможно представить число в римских цифрах")
	}
	result := ""
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			result += symbols[i]
			num -= values[i]
		}
	}
	return result
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("Деление на ноль невозможно")
		}
		return a / b
	default:
		panic("Неизвестная операция")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например: 3 + 5 или IV * II):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		panic("Неверный формат ввода. Ожидается: a + b")
	}

	a, operator, b := parts[0], parts[1], parts[2]

	// Проверка, арабские или римские числа
	isRomanMode := isRoman(a) && isRoman(b)
	isArabicMode := !isRoman(a) && !isRoman(b)

	if !isRomanMode && !isArabicMode {
		panic("Нельзя смешивать арабские и римские числа")
	}

	var num1, num2 int
	if isRomanMode {
		num1 = romanToInt(a)
		num2 = romanToInt(b)
	} else {
		num1, _ = strconv.Atoi(a)
		num2, _ = strconv.Atoi(b)
	}

	// Проверка на диапазон от 1 до 10
	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		panic("Числа должны быть в диапазоне от 1 до 10")
	}

	// Выполнение операции
	result := calculate(num1, num2, operator)

	// Вывод результата в зависимости от режима
	if isRomanMode {
		if result < 1 {
			panic("Результат римских чисел не может быть меньше I")
		}
		fmt.Println("Результат:", intToRoman(result))
	} else {
		fmt.Println("Результат:", result)
	}
}
