package main

import (
	"fmt"
	"github.com/recktt77/bankaccount-exe4.git/bancking"
	"strconv"
)

func main() {
	bank := bancking.Bank{}
	bank.Init()
	function := ""
	name := ""
	var balance float64
	var id string
	for {
		println("Choose the function (addBankAccount, Deposit, Withdraw, GetBalance, Transaction, or exit)")
		fmt.Scan(&function)

		switch function {
		case "addBankAccount":
			fmt.Println("Enter account name and balance")
			fmt.Scan(&name, &balance)
			if name == "" || balance < 0 {
				fmt.Println("cannot be empty or negative number. ")
				continue
			}
			bank.AddAcount(name, balance)
		case "Deposit":
			fmt.Println("Enter the ID and amount of money")
			fmt.Scan(&id, &balance)
			if id == "" || balance < 0 {
				fmt.Println("cannot be or negative number")
				continue
			}
			bank.Deposit(id, balance)
		case "Withdraw":
			fmt.Println("Enter the ID and amount of money")
			fmt.Scan(&id, &balance)
			if id == "" || balance < 0 {
				fmt.Println("cannot be or negative number")
				continue
			}
			bank.Withdraw(id, balance)
		case "GetBalance":
			fmt.Println("Enter the ID")
			fmt.Scan(&id)
			bank.GetBalance(id)
		case "Transaction":
			var transactions []float64
			fmt.Println("Enter the ID:")
			fmt.Scan(&id)
			fmt.Println("Enter transaction amounts (positive for deposit, negative for withdrawal). Type 'done' to finish:")

			for {
				var input string
				fmt.Scan(&input)
				if input == "done" {
					break
				}
				amount, err := strconv.ParseFloat(input, 64)
				if err != nil {
					fmt.Println("Invalid input, please enter a number or 'done'.")
					continue
				}
				transactions = append(transactions, amount)
			}

			bank.Transaction(id, transactions)
		case "exit":
			fmt.Println("Turned off")
			return
		default:
			fmt.Println("Sometime keep using something usual is better then something new")
		}
	}
}
