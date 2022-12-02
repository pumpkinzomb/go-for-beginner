package accounts

import "errors"

type account struct {
	name string
	balance int
}

func NewAccount(name string) *account {
	account := account {
		name: name,
		balance: 0,
	}
	return &account
}

func (a *account) Deposit(amount int) {
	a.balance += amount
}

func (a *account) Withdraw(amount int) error {
	if(a.balance > amount){
		a.balance -= amount;
		return nil
	}
	return errors.New("you have not enough balances.")
}

func (a account) Balances() int {
	return a.balance
}