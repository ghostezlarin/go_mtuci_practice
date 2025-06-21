package main

import "fmt"

// Определение структуры Student
type Student struct {
	Name   string
	Age    int
	Course int
	GPA    float64
}

// Функция для создания нового студента
func NewStudent(name string, age, course int, gpa float64) Student {
	return Student{
		Name:   name,
		Age:    age,
		Course: course,
		GPA:    gpa,
	}
}

// Функция для вывода информации о студенте
func (s Student) PrintInfo() {
	fmt.Printf("Студент: %s\n", s.Name)
	fmt.Printf("Возраст: %d\n", s.Age)
	fmt.Printf("Курс: %d\n", s.Course)
	fmt.Printf("Средний балл: %.2f\n", s.GPA)
}

// Функция для повышения курса
func (s *Student) Promote() {
	s.Course++
	fmt.Printf("%s переведен на %d курс\n", s.Name, s.Course)
}

// Функция для обновления среднего балла
func (s *Student) UpdateGPA(newGPA float64) {
	s.GPA = newGPA
	fmt.Printf("Средний балл %s обновлен: %.2f\n", s.Name, s.GPA)
}

func main() {
	// Создаем студента
	student := NewStudent("Иван Иванов", 20, 2, 4.5)

	// Выводим информацию
	student.PrintInfo()

	// Повышаем курс
	student.Promote()

	// Обновляем средний балл
	student.UpdateGPA(4.7)

	// Выводим обновленную информацию
	fmt.Println("\nОбновленные данные:")
	student.PrintInfo()
}
