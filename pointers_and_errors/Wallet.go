package main

import (
	"errors"
	"fmt"
)

type IguanaDolla int

type Wallet struct {
	balance IguanaDolla
}

type Stringer interface {
	String() Stringer
}

func (w *Wallet) Withdraw(amount IguanaDolla) error {
	if amount > w.balance {
		return errors.New("cannont withdraw, insufficient funds")
	} else {
		w.balance -= amount
	}
	return nil
}

func (w *Wallet) Deposit(amount IguanaDolla) {
	w.balance += amount
}

func (w *Wallet) Balance() IguanaDolla {
	return w.balance
}

func (i IguanaDolla) String() string {
	return fmt.Sprintf("%d ID", i)
}
