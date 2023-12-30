package blogposts

import (

	"testing"
	"fmt"
	"reflect"
)	
	
func TestSliceAndArray(t *testing.T){

	t.Run(" Test slice " , func(t *testing.T){

		a := []int { 1 , 2 , 3 , 4}
		var b  []int
		c := make([]int , 4)
		d := make([]int , 4 , 6)
		fmt.Printf(" a len %d , capacity %d  , data %v \n" , len(a) , cap(a) , a)
		fmt.Printf(" b len %d , capacity %d  , data %v \n" , len(b) , cap(b) , b)
		fmt.Printf(" c len %d , capacity %d  , data %v \n" , len(c) , cap(c) , c)
		fmt.Printf(" d len %d , capacity %d  , data %v \n" , len(d) , cap(d) , d)

		e := a

		a[1] = 99
		if !reflect.DeepEqual(a , e){
			t.Error(" they should be same , assign by reference")
		}

		c =append(c , 5)
		fmt.Printf(" c after append len %d , capacity %d  , data %v \n" , len(c) , cap(c) , c)
		c =append(c , 6)
		fmt.Printf(" c after append len %d , capacity %d  , data %v \n" , len(c) , cap(c) , c)


		var arr [4]int = [4]int{1, 2, 3, 4}

		s := arr[:3]
		fmt.Printf( " s len %d capacity %d , data %v\n" , len(s) , cap(s) ,s )

		s = arr[1:]
		fmt.Printf( " s len %d capacity %d , data %v\n" , len(s) , cap(s) ,s )
		s = arr[1:3]
		fmt.Printf( " s len %d capacity %d , data %v\n" , len(s) , cap(s) ,s )


		// at this time slice is on arr
		s[0] = 6
		fmt.Printf( " round 1 s len %d capacity %d , data %v\n" , len(s) , cap(s) ,s )
		fmt.Printf( " round 1 arr len %d capacity %d , data %v\n" , len(arr) , cap(arr) ,arr )

		// still update on original array
		s = append(s , 8)
		fmt.Printf( " round 2 s append len %d capacity %d , data %v\n" , len(s) , cap(s) ,s )
		fmt.Printf( " round 2 arr len %d capacity %d , data %v\n" , len(arr) , cap(arr) ,arr )

		// at this time the slice run out of space , so append creats a new one
		s = append(s , 9)
		fmt.Printf( " round 3 s append len %d capacity %d , data %v\n" , len(s) , cap(s) ,s )
		fmt.Printf( " round 3 arr len %d capacity %d , data %v\n" , len(arr) , cap(arr) ,arr )

	})

	t.Run(" Array aissign by value" , func ( t *testing.T){
		a := [4]int { 1 , 2 , 3 , 4}
		b := a
	
		b[1] = 999
	
		fmt.Printf(" a len %d , capacity %d  , data %v \n" , len(a) , cap(a) , a)
		fmt.Printf(" b len %d , capacity %d  , data %v \n" , len(b) , cap(b) , b)
	
		if reflect.DeepEqual(a , b){
			t.Error(" they should be different , assign by value")
		}

	})
}

func TestArray(t *testing.T){

}