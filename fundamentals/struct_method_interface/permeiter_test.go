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

func TestAreaFunc( t  * testing.T){

	rectangle :=Rectangle{12.0 , 6.0}
	got := Area(rectangle)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

}

func TestArea(t *testing.T){

	checkArea := func(t testing.TB , shape Shape , want float64){
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf(" got %g wang %g" , got , want)
		}
	}

	t.Run("rectangle " , func(t *testing.T){
		rectangle := Rectangle{12.0 , 6.0}
		checkArea(t , rectangle , 72.0)
	})

	t.Run("Circle" , func(t *testing.T){
		circle := Circle{10}
		checkArea(t , circle , math.Pi * 100.0)
	})

}

func TestArea2( t *testing.T){

	areaTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{12.0  , 6.0} , 72.0},
		{Circle{10.0} , math.Pi*100},
		{shape:Triangle{Base:12, Height:6}, want:36.0},
	}

	for _ , tt := range areaTests{
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf(" got %g wang %g" , got , tt.want)
		}
	}
}