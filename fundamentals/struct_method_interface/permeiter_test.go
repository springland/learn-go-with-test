package struct_method_interface

import (
	"testing"
	"math"
)


func TestPermiter( t *testing.T){

	t.Run(" test rectangle permiter" , func( t *testing.T){

		rectangle :=Rectangle{10.0 , 10.0}
		got := rectangle.Perimeter()
		want := 40.0
	
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	
	})

	t.Run(" circle permiter" , func(t *testing.T){

		circle := Circle{10.0}
		got := circle.Perimeter()
		want :=  math.Pi * 10.0*2

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}

	})
}

func TestArea( t  * testing.T){

	rectangle :=Rectangle{12.0 , 6.0}
	got := Area(rectangle)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

}