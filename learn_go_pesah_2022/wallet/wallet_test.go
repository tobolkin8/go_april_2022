package wallet

import (
	"hello/tests"
	"testing"
)

func TestDeposit(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))
        got := wallet.Balance()
        want := Bitcoin(10)

        tests.FailedTestResults(t,got,want)
								tests.CorrectResults(got,want)
}
func TestWithdraw(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(20)}
									err:= wallet.Withdraw(Bitcoin(10))
									got:= wallet.Balance()
									want := Bitcoin(10)

        tests.FailedTestResults(t,got,want)
								tests.ErrorResults(t,err)
								tests.CorrectResults(got,want)
}

func TestWithdrawNoMoney(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(20)}
									err:= wallet.Withdraw(Bitcoin(50))
								tests.ErrorResults(t,err)
}

