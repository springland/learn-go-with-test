package integers

import (
	"testing" ;
	"fmt"
)
func TestAdd(t *testing.T) {

	t.Run( " test add" , func( t *testing.T){
		sum := Add(2 , 2)
		expected := 4
		assertEquals(sum , expected , t)
	})
}

func assertEquals( actual int , expected int , t *testing.T){

	if(actual != expected){
		t.Errorf(" Acutal is %d , expected is %d" , actual , expected)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}