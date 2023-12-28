package property_based

import (
	"strings"

)
var  numberToRomanMap = map[int]string {
	1000 : "M",
	500: "D",
	100: "C",
	50 : "L",
	10: "X",
	5: "V",
	1: "I",
}

func ConvertToRoman( arabic  int) string{


	return numberToRoman(arabic , 1)
	// result := ""

	// switch arabic {
	// 	case 1 , 2 , 3:
	// 		for index:= 0 ; index < arabic ; index++{
	// 			result += "I"
	// 		}	
	// 	case 4:
	// 		result = "IV"
	// 	case 5:
	// 		result = "V"
	// 	case 6 , 7 , 8:
	// 		result = "V"
	// 		for index := 5 ; index < arabic ; index ++{
	// 			result += "I"
	// 		}
	// 	case 9:
	// 		result = "IX"	
	// }
	// return result
}

func numberToRoman( arabic int , base int) string{
	result := ""  

	for arabic >0 {

		remainder := arabic % 10
		result = 	singleDigitNumberToRoman(remainder , base)+ result 
		base = base *10
		arabic = arabic/10
	}

	return result
} 

func singleDigitNumberToRoman(arabic int , base int) string {
	var result  strings.Builder

	switch arabic {
		case 0:
			result.WriteString("")
		case 1 , 2 , 3:
			for index:= 0 ; index < arabic ; index++{
				result.WriteString(numberToRomanMap[base ]) 
			}	
		case 4:
			result.WriteString(numberToRomanMap[base] + numberToRomanMap[5*base])
			
		case 5:
			result.WriteString(numberToRomanMap[5*base])
		case 6 , 7 , 8:
			result.WriteString(numberToRomanMap[base*5])	
			for index:=5 ; index < arabic ; index++ {
				result.WriteString(numberToRomanMap[base ])
			}			
			
		case 9: 
			result.WriteString(numberToRomanMap[base] + numberToRomanMap[base*10])

	}
	return result.String()				
}