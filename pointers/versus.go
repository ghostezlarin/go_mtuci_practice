package main

import "fmt"

// Структура для демонстрации
type Person struct {
	Name string
	Age  int
}

// Функция с передачей по значению (копия структуры)
func incrementAgeByValue(p Person) {
	p.Age++
	fmt.Println("Внутри incrementAgeByValue:", p.Age)
}

// Функция с передачей по указателю (оригинал структуры)
func incrementAgeByPointer(p *Person) {
	p.Age++
	fmt.Println("Внутри incrementAgeByPointer:", p.Age)
}

func main() {
	person := Person{"Алексей", 30}

	fmt.Println("1. Передача по значению:")
	fmt.Println("До вызова:", person.Age)
	incrementAgeByValue(person)
	fmt.Println("После вызова:", person.Age)

	fmt.Println("\n2. Передача по указателю:")
	fmt.Println("До вызова:", person.Age)
	incrementAgeByPointer(&person)
	fmt.Println("После вызова:", person.Age)
}
