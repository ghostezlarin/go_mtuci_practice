package main

import "fmt"

func main() {
	var a, b float64

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	// Выполнение и вывод всех арифметических операций
	fmt.Printf("\nРезультаты операций:\n")
	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
	fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)

	// Особый случай для деления
	if b != 0 {
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
	} else {
		fmt.Printf("%.2f / %.2f = ошибка (деление на ноль)\n", a, b)
	}

	// Операция остатка от деления (только для целых чисел)
	if b != 0 {
		fmt.Printf("%d %% %d = %d\n", int(a), int(b), int(a)%int(b))
	} else {
		fmt.Printf("%d %% %d = ошибка (деление на ноль)\n", int(a), int(b))
	}
}
