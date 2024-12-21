package bancking

import (
	"fmt"
	"strconv"
)

type UserAccount interface {
	Deposit(string, float64)
	Withdraw(string, float64)
	GetBalance(string)
}
type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}
type Bank struct {
	BankAccounts map[string]*BankAccount
	nextId       int
}

func (ba *Bank) Init() {
	if ba.BankAccounts == nil {
		ba.BankAccounts = make(map[string]*BankAccount)
	}
}
func (ba *Bank) AddAcount(HolderName string, Balance float64) {
	id := strconv.Itoa(ba.nextId)
	ba.BankAccounts[id] = &BankAccount{
		AccountNumber: id,
		HolderName:    HolderName,
		Balance:       Balance,
	}
	ba.nextId++
	fmt.Printf("Account '%s' added", HolderName)
}
func (ba *Bank) Deposit(id string, amount float64) {
	acount, ok := ba.BankAccounts[id]
	if !ok {
		fmt.Println("Account doesn't exist.")
		return
	}
	acount.Balance += amount
	fmt.Printf("balance added: '%v' ", acount.Balance)
}
func (ba *Bank) Withdraw(id string, amount float64) {
	acount, ok := ba.BankAccounts[id]
	if !ok {
		fmt.Println("Account doesn't exist.")
		return
	}
	if acount.Balance < amount {
		fmt.Printf("Imposible you have only: '%v' ", acount.Balance)
		return
	}
	acount.Balance -= amount
	fmt.Printf("balance withdrawed: '%v' ", acount.Balance)
}
func (ba *Bank) GetBalance(id string) {
	account, ok := ba.BankAccounts[id]
	if !ok {
		fmt.Println("Account doesn't exist.")
		return
	}
	fmt.Printf("you have on balance: '%v' ", account.Balance)
}

func (ba *Bank) Transaction(id string, transactions []float64) {
	account, ok := ba.BankAccounts[id]
	if !ok {
		fmt.Println("Account doesn't exist.")
		return
	}
	for _, t := range transactions {
		if t > 0 {
			account.Balance += t
		} else if account.Balance >= -t {
			account.Balance += t
		} else {
			fmt.Printf("Insufficient funds for transaction: %v\n", t)
		}
	}
	fmt.Printf("Final balance for account '%s': %.2f\n", id, account.Balance)
}
