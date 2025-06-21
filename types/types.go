package main

import "fmt"

func main() {
	// Целочисленные типы
	var int8Var int8 = -128
	var int16Var int16 = -32768
	var int32Var int32 = -2147483648
	var int64Var int64 = -9223372036854775808
	var intVar int = 2147483647 // зависит от платформы (32 или 64 бита)

	var uint8Var uint8 = 255
	var uint16Var uint16 = 65535
	var uint32Var uint32 = 4294967295
	var uint64Var uint64 = 18446744073709551615
	var uintVar uint = 4294967295 // зависит от платформы

	// Числа с плавающей точкой
	var float32Var float32 = 3.1415926535
	var float64Var float64 = 3.141592653589793

	// Комплексные числа
	var complex64Var complex64 = 3 + 4i
	var complex128Var complex128 = 3 + 4i

	// Логический тип
	var boolVar bool = true

	// Строки
	var stringVar string = "Hello, Go!"

	// Символы (псевдоним int32 для Unicode)
	var runeVar rune = 'Я'

	// Байт (псевдоним uint8)
	var byteVar byte = 255

	fmt.Println("Целочисленные (со знаком):")
	fmt.Println("int8:", int8Var, "int16:", int16Var, "int32:", int32Var, "int64:", int64Var, "int:", intVar)

	fmt.Println("\nЦелочисленные (без знака):")
	fmt.Println("uint8:", uint8Var, "uint16:", uint16Var, "uint32:", uint32Var, "uint64:", uint64Var, "uint:", uintVar)

	fmt.Println("\nЧисла с плавающей точкой:")
	fmt.Println("float32:", float32Var, "float64:", float64Var)

	fmt.Println("\nКомплексные числа:")
	fmt.Println("complex64:", complex64Var, "complex128:", complex128Var)

	fmt.Println("\nПрочие типы:")
	fmt.Println("bool:", boolVar, "string:", stringVar, "rune:", runeVar, "byte:", byteVar)
}
