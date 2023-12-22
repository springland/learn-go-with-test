package pointer_error

import(
	"fmt"
	"errors"
) 
type BitCoin int

type Wallet struct {
	balance BitCoin
}


func (wallet Wallet) Balance() BitCoin {
	return wallet.balance
}

func (wallet *Wallet) Deposite(amount BitCoin) {
	wallet.balance += amount 
}

func (wallet *Wallet) Withdraw(amount BitCoin) error {

	if amount > wallet.balance {
		return errors.New("oh no , there is not enough balance")
	}
	wallet.balance -= amount 

	return nil
}


type Stringer interface {
	String() string
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC" , b)
}