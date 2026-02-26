// Package pointerr pointers & errors
package pointerr

import "fmt"

type Bitcoin int

// String implementa la interfaz Stringer
// permite definir como se escribe el tipo (Bitcoin) cuando se usa %s
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// TODO: refactor assertError...

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return fmt.Errorf("no insufficient coin: %v", amount)
	}
	w.balance -= amount
	return nil
}
