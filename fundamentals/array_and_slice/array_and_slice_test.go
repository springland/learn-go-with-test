package array_and_slice

import ( 
	"testing";
	"reflect";
	"fmt"
)
func TestSumArrayByIndex(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := SumArrayByIndex(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumArrayByRange( t *testing.T){

	numbers := [5]int{1, 2, 3, 4, 5}

	got := SumArrayByRange(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}


}

func  TestSumSliceByRange( t *testing.T){
	numbers := []int{1 , 2 , 3 , 4 , 5}

	got := SumSliceByRange(numbers)
	want := 15
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}


}

func TestSumAll( t *testing.T){
	got := SumAll( []int {1 , 2} , [] int{0 ,9})
	want := []int{3 , 9}

	if !reflect.DeepEqual(got, want)  {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllAppend( t* testing.T){
	got := SumAllAppend( []int {1 , 2} , [] int{0 ,9})
	want := []int{3 , 9}

	if !reflect.DeepEqual(got, want)  {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestSumALlTails( t* testing.T){
	t.Run(" Sum All Tails" , func(t *testing.T){
		got := SumAllTails( []int{ 1 , 2 , 3} , []int { 4 , 5, 6})
		want := []int{ 5 , 11}

		if !reflect.DeepEqual(got, want)  {
			t.Errorf("got %v want %v", got, want)
		}
	

	})


	t.Run(" Sum All Tails with empty" , func(t *testing.T){
		got := SumAllTails( []int{ } , []int { 4 , 5, 6})
		want := []int{ 0 , 11}

		if !reflect.DeepEqual(got, want)  {
			t.Errorf("got %v want %v", got, want)
		}
	

	})

}

func TestSlice(t *testing.T){

	array := [5]int{ 1 , 2  ,3 , 4 , 5}
	slice := array[2:4]

	fmt.Printf(" %v" , slice)

}