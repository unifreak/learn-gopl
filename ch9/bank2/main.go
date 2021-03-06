package bank

var (
	sema = make(chan struct{}, 1) // a binary semaphor guarding balance
	balance int
)

func Deposit(amoutn int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<- sema // release token
	return b
}