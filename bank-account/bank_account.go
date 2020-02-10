package account

import (
	"sync"
)

// Account Specifies the object Account
type Account struct {
	sync.Mutex
	balance int64
	closed  bool
}

// Close the account
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}

	payout = a.balance
	a.closed = true

	return payout, true
}

// Balance Checks the balance.
func (a *Account) Balance() (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}

	return a.balance, true

}

// Deposit deposits amount, accepts an argument amount specifying amount to deposit.
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		return 0, false
	}

	if a.balance+amount < 0 {
		return 0, false
	}

	a.balance += amount

	return a.balance, true
}

// Open Opens account with specified initial amount.
func Open(initialDeposit int64) *Account {
	if initialDeposit >= 0 {
		return &Account{balance: initialDeposit, closed: false}
	}
	return nil
}
