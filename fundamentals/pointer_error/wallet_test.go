package pointer_error
import "testing"

func TestWallet(t *testing.T){

	t.Run("Deposite " , func(t *testing.T){

		wallet := Wallet {}

		wallet.Deposite(BitCoin(10))
		
		got := wallet.Balance()
	
		want := BitCoin(10)
	
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	
	})

	t.Run("Withdraw" , func( t *testing.T) {
		wallet := Wallet {BitCoin(10)}

		
		wallet.Withdraw(BitCoin(5))
		got := wallet.Balance()
	
		want := BitCoin(5)
	
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}


	})


	t.Run("WithdrawExceedBalance" , func( t *testing.T){

		wallet := Wallet {BitCoin(10)}

		
		err := wallet.Withdraw(BitCoin(15))

		if  err == nil{
			t.Error("Wanted an error but did get one")
		}
	})

}