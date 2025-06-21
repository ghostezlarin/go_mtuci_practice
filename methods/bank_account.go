package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	owner   string
	balance float64
}

// Конструктор нового счета
func NewBankAccount(owner string, initialBalance float64) *BankAccount {
	return &BankAccount{
		owner:   owner,
		balance: initialBalance,
	}
}

// Метод пополнения счета
func (acc *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	acc.balance += amount
	return nil
}

// Метод снятия средств
func (acc *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	if amount > acc.balance {
		return errors.New("недостаточно средств")
	}
	acc.balance -= amount
	return nil
}

// Метод проверки баланса
func (acc *BankAccount) Balance() float64 {
	return acc.balance
}

// Метод для вывода информации
func (acc *BankAccount) Info() {
	fmt.Printf("Владелец счета: %s\n", acc.owner)
	fmt.Printf("Текущий баланс: %.2f руб.\n", acc.balance)
}

func main() {
	// Создаем новый счет
	account := NewBankAccount("Алексей Иванов", 1000.0)
	account.Info()

	// Пополняем счет
	err := account.Deposit(500.0)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("\nПосле пополнения на 500 руб.")
		account.Info()
	}

	// Снимаем средства
	err = account.Withdraw(200.0)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("\nПосле снятия 200 руб.")
		account.Info()
	}

	// Пытаемся снять слишком много
	err = account.Withdraw(2000.0)
	if err != nil {
		fmt.Println("\nОшибка при снятии:", err)
	}
}
