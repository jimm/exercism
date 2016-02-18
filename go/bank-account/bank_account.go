package account

import "sync"

type Account struct {
	sync.Mutex
	open    bool
	balance int64
}

func Open(initalDeposit int64) *Account {
	if initalDeposit < 0 {
		return nil
	}
	return &Account{open: true, balance: initalDeposit}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	ok = a.open
	payout = a.balance
	a.open = false
	a.balance = 0
	return
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	ok = a.open
	balance = a.balance
	return
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	ok = true
	if !a.open {
		ok = false
	} else {
		newBalance = a.balance + amount
		if newBalance < 0 {
			ok = false
		} else {
			a.balance = newBalance
		}
	}
	return
}
