package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds") //var keyword signfies the variable as global to the package
//also errors are values in go so we can encapsulate them into variables like done here!

type Bitcoin int

type Stringer interface { //This interface is defined in the fmt package and lets you define how your type is printed when used with the %s format string in prints.
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

//dereferencing in Go is automatic

// If you don't make the method receiever a pointer then whenever you call it, it will make a copy of the wallet and not directly affect the variables you want them to
func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds // creates a new error with a message of my choosing, which in this case is "oh no"
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
