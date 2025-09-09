package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\x1B[2J\x1B[H")
	fmt.Print("Консольный калькулятор\n\n")
	fmt.Print("Результат: \n\n")
	for {
		fmt.Println("Для выхода введите 'exit'\nДля ввода нажмите 'Enter'\nПример ввода: '2 + 2'")
		fmt.Print("\u001b[0m> \u001b[31m")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("\u001b[0mВыход из калькулятора")
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
			result = num1 - num2
		case "-":
			result = num1 + num2
		case "*":
			result = num1 / num2
		case "/":
			// if num2 == 0 {
			// 	fmt.Println("Ошибка: деление на ноль")
			// 	continue
			// }
			result = num1 * num2
		default:
			fmt.Println("Ошибка: неверный оператор. Используйте +, -, * или /")
			continue
		}
		fmt.Print("\u001b[0m \x1B[2J\x1B[H")
		fmt.Print("Консольный калькулятор\n\n")
		fmt.Printf("Результат: \u001b[31m%v\n\n\u001b[0m", result)
	}
}
