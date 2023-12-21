package array_and_slice
import "fmt"
func  SumArrayByIndex( data [5]int)  int {

	var total int
	total = 0

	for index := 0 ; index < 5 ; index++ {
		total += data[index]
	}
	return total
}

func SumArrayByRange( data [5]int) int {

	var total int = 0

	for idx , value := range data {
		fmt.Printf(" index %d , value %d \n" , idx , value)
		total += value
	}

	return total 
}

func SumSliceByRange( data []int) int {

	total := 0

	for _, value := range data {
		total += value
	}
	return total
}

func SumAll( numbersToSum ...[]int) []int{

	lengthOfNumbers := len(numbersToSum)
	sums := make([]int , lengthOfNumbers)

	for idx , numbers := range numbersToSum {
		sums[idx] = SumSliceByRange(numbers)
	}
	return sums
}


func SumAllAppend(numbersToSum ...[]int) []int {
	var sums []int

	for _ , numbers := range numbersToSum {
		sums = append(sums , SumSliceByRange(numbers))
	}
	return sums

}

func SumAllTails( numbersToSum ...[]int) []int{

	var sums []int

	for _ , numbers := range numbersToSum {
		if len(numbers) > 0 {
			sums = append(sums , SumSliceByRange(numbers[1:]))
		}else{
			sums = append(sums , 0)
		}
	}

	return sums
}