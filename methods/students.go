package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthYear int // год рождения
	Course    int
	GPA       float64 // средний балл (0-5)
}

// Метод вычисления возраста
func (s Student) Age() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

// Метод получения статуса
func (s Student) Status() string {
	switch {
	case s.GPA >= 4.5:
		return "отличник"
	case s.GPA >= 3.5:
		return "хорошист"
	case s.GPA >= 2.0:
		return "троечник"
	default:
		return "двоечник"
	}
}

// Метод для вывода информации
func (s Student) Info() {
	fmt.Printf("Студент: %s\n", s.Name)
	fmt.Printf("Возраст: %d лет\n", s.Age())
	fmt.Printf("Курс: %d\n", s.Course)
	fmt.Printf("Средний балл: %.1f (%s)\n", s.GPA, s.Status())
}

func main() {
	student := Student{
		Name:      "Иван Петров",
		BirthYear: 2000,
		Course:    3,
		GPA:       4.7,
	}

	student.Info()
}
