package reflection

import(
	"reflect"

) 

func walk(x interface{} , fn func(input string)) {

	val := getValue(x)

	switch val.Kind() {
		case reflect.String:
			
			fn(val.String())
		case reflect.Struct:

			for i:=0 ;i < val.NumField() ; i++{
				walk( val.Field(i).Interface() , fn)
			}
		case reflect.Slice:
			for i:= 0 ; i < val.Len() ; i++{
				walk(val.Index(i).Interface() , fn)
			}	

		case reflect.Array:
			for i:= 0 ; i < val.Len() ; i++{
				walk(val.Index(i).Interface() , fn)
			}	

		case reflect.Map:
			itr := val.MapRange()
			for itr.Next(){
				walk(itr.Value().Interface()  , fn)
			}	
	}

}


func walk2(x interface{} , fn func(input string)) {

	val := getValue(x)
	var getField func(int) reflect.Value
	numberOfValues := 0

	switch val.Kind() {
		case reflect.String:
			
			fn(val.String())
		case reflect.Struct:
			numberOfValues = val.NumField()
			getField = val.Field

		case reflect.Slice:

			numberOfValues = val.Len()
			getField = val.Index


		case reflect.Array:
			numberOfValues = val.Len()
			getField = val.Index

	}

	for i:=0 ;i < numberOfValues ; i++{
		walk( getField(i).Interface() , fn)
	}

}


func getValue(x interface{}  ) reflect.Value {

	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val

}