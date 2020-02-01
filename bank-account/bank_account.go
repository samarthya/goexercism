package account

import (
	"fmt"
	"sync"
)

// Banking interface
type Banking interface {
	Close() (int64, bool)
	Balance() (int64, bool)
	Deposit() (int64, bool)
}

// Account type
type Account struct {
	balance int64
	lock    sync.Mutex
	state   bool
}

// Close the account
func (my *Account) Close() (payout int64, ok bool) {
	payout = 0
	ok = true
	my.lock.Lock()
	if (my.state) == true {
		fmt.Println(" Close the account with balance", my.balance)
		payout = my.balance
		my.state = false
	} else {
		ok = false
	}
	my.lock.Unlock()
	return payout, ok
}

// Balance - Checks the balance.
func (my *Account) Balance() (balance int64, ok bool) {
	balance = 0
	ok = false
	my.lock.Lock()
	if (my.state) == true {
		balance = my.balance
		fmt.Println(" Account with balance", my.balance)
		ok = true
	}
	my.lock.Unlock()
	return balance, ok
}

// Deposit deposits amount.
func (my *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	ok = false
	my.lock.Lock()
	if (my.state) == true {
		newBalance = my.balance
		if (amount + my.balance) >= 0 {
			my.balance += amount
			newBalance = my.balance
			fmt.Println(" Account deposit ", amount, newBalance)
			ok = true
		}
	}
	my.lock.Unlock()
	return newBalance, ok
}

// Open Opens account
func Open(initialDeposit int64) *Account {
	if initialDeposit >= 0 {
		fmt.Println(" Open the account with balance", initialDeposit)
		return &Account{balance: initialDeposit, state: true}
	}
	return nil
}
