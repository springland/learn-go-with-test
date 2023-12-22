package mymap
import "testing"

func TestSearch(t *testing.T){
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run(" Test known word" , func( t *testing.T){
		got , _ := dictionary.Search( "test")
		want   := "this is just a test"
	
		assertStrings(t , got , want)

	})

	t.Run(" Test unknown word" , func( t *testing.T){
		_ , err := dictionary.Search( "unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t , err.Error() , want)
	})

	
}




func TestAdd( t *testing.T){
	

	t.Run(" Add no exist word" , func ( t *testing.T){
		dict := Dictionary{}
		dict.Add("test", "this is just a test")
		want :=  "this is just a test"
		got , err := dict.Search("test")
		if(err != nil){
			t.Fatal("should find added word:", err)
		}
		assertStrings(t , got , want)
	

	})

	t.Run(" Add Existing word" , func( t* testing.T){

		dict := Dictionary{"test":"value"}
		err :=dict.Add("test", "this is just a test")
		want :=  "cannot add word because it already exists"
		if(err == nil){
			t.Fatal("should find added word:", err)
		}
		assertStrings(t , err.Error() , want)

	})

}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}