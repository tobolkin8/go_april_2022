package wallet

import (
	"errors"
	"fmt"
	"hello/str_consts"
)

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}
type Stringer interface {
    String() string
}

func (w *Wallet) Deposit(amount Bitcoin){
	w.balance += amount
}

func (w *Wallet) Balance() (Bitcoin) {
	return w.balance
}

func (b Bitcoin) String() string {
    return fmt.Sprintf(str_consts.WALLET_BTC_FORMAT, b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	 if amount > w.balance {
        return errors.New(str_consts.WALLET_NO_FOUNDS)
    }
	w.balance -= amount
	return nil
}