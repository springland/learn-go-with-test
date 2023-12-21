package struct_method_interface
import "math"
type Rectangle struct {
	width float64
	height float64
}


func ( rect Rectangle) Perimeter ( ) float64{

	return (rect.width + rect.height)*2
}

func  Area( rect Rectangle) float64 {
	return rect.width * rect.height
}


type Circle struct {
	radius float64
}


func (circle Circle) Perimeter() float64 {
	return math.Pi * circle.radius*2
}


