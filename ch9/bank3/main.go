package bank

import "sync"

var (
	// By convention, the variables guarded by a mutex are declared immediately after
	// the declaration of the mutex itself.
	mu sync.Mutext // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}