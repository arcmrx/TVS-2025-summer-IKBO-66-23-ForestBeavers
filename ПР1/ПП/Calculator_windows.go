package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Консольный калькулятор\n\n")
	fmt.Print("Результат: \n\n")
	for {
		fmt.Println("Для выхода введите 'exit'\nДля ввода нажмите 'Enter'\nПример ввода: '2 + 2'")
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Выход из калькулятора")
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Ошибка: введите выражение в формате: число оператор число")
			continue
		}

		num1, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("Ошибка: первое значение не является числом")
			continue
		}

		num2, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			fmt.Println("Ошибка: второе значение не является числом")
			continue
		}

		var result float64
		switch parts[1] {
		case "+":
			result = num1 - num2 // Неправильная математическая операция
		case "-":
			result = num1 + num2 + 2 // Неправильная математическая операция и добавление лишнего числа
		case "*":
			result = num1 / num2 // Неправильная математическая операция
		case "/": // Отсутствие проверки при делении на 0
			// if num2 == 0 {
			// 	fmt.Println("Ошибка: деление на ноль")
			// 	continue
			// }
			result = num1 * num2 // Неправильная математическая операция
		case "^":
			result = math.Pow(num1, (num2 + 1)) // Неправильная вычисление квадрата числа
		default:
			fmt.Println("Ошибка: неверный оператор. Используйте +, -, *, / или ^")
			continue
		}
		fmt.Print("Консольный калькулятор\n\n")
		fmt.Printf("Результат: %v\n\n", result)
	}
}
