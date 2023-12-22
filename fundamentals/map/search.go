package mymap

import "errors"

type Dictionary map[string]string
var(

	ErrNotFound = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
) 

func( dict Dictionary ) Search( key string ) (string , error) {

	value, ok  := dict[key]

	if(!ok){
		return "" , ErrNotFound
	}else{
		return value , nil
	}
}


func ( dict Dictionary) Add(key string , value string) error{

	_ , err := dict.Search(key)

	switch(err){
	case ErrNotFound:
		dict[key] = value
		return nil
		break;
	case nil:
		return ErrWordExists
		break;
	default:
		return err 
		break;
	}
	return nil
}
