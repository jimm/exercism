package account

import "sync"

type Account struct {
	mutex   *sync.Mutex
	open    bool
	balance int64
}

func Open(initalDeposit int64) *Account {
	if initalDeposit < 0 {
		return nil
	}
	return &Account{&sync.Mutex{}, true, initalDeposit}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.lock()
	ok = a.open
	payout = a.balance
	a.open = false
	a.balance = 0
	a.unlock()
	return
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.lock()
	ok = a.open
	balance = a.balance
	a.unlock()
	return
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	ok = true
	a.lock()
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
	a.unlock()

	return
}

func (a *Account) lock() {
	a.mutex.Lock()
}

func (a *Account) unlock() {
	a.mutex.Unlock()
}
