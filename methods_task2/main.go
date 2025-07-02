package main

import "fmt"

type BankAccount struct {
	owner   string
	balance float64
}

func (a *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Сумма пополнения должна быть положительной.")
		return
	}
	a.balance += amount
	fmt.Printf("Счет пополнен на %.2f. Новый баланс: %.2f\n", amount, a.balance)
}

func (a *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Сумма снятия должна быть положительной.")
		return
	}
	if a.balance < amount {
		fmt.Printf("Недостаточно средств. Баланс: %.2f, попытка снять: %.2f\n", a.balance, amount)
		return
	}
	a.balance -= amount
	fmt.Printf("Со счета снято %.2f. Новый баланс: %.2f\n", amount, a.balance)
}

func (a BankAccount) CheckBalance() {
	fmt.Printf("Баланс счета владельца %s: %.2f\n", a.owner, a.balance)
}

func main() {
	account := BankAccount{owner: "Иванов И.И.", balance: 1000.0}
	account.CheckBalance()
	fmt.Println("---")

	account.Deposit(500.50)
	account.Withdraw(200.0)
	fmt.Println("---")

	account.Withdraw(1500.0)
	fmt.Println("---")

	account.CheckBalance()
}
