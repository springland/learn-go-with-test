package main

import (
 "testing";
 //"fmt"	
)
func TestHelloWorld( t *testing.T) {
	got := HelloWorld()

	want := "Hello World"

	//fmt.Printf("got %q  want %q" , got , want)
	if( got != want){
		t.Errorf( "got %q want %q" , got ,  want)
	}
}

func TestHello( t *testing.T){

	got := Hello("Chris" , "")
	want := "Hello Chris"
	if( got != want){
		t.Errorf(" got %q want %q" , got , want)
	}

	t.Run("Say Hello to People , default language" , func(t *testing.T){

		got := Hello("Chris" , "")
		want := "Hello Chris"
		assertEquals(got , want , t)	
	})

	t.Run("Say Hello to World , default language" , func(t *testing.T){

		got := Hello("" , "")
		want := "Hello World"
		assertEquals(got , want , t)
	})

	t.Run("Say Hello to People , Spanish" , func(t *testing.T){

		got := Hello("Chris" , "spanish")
		want := "Hola Chris"
		assertEquals(got , want , t)	

	})

	t.Run("Say Hello to People , French" , func(t *testing.T){

		got := Hello("Chris" , "french")
		want := "Bonjour Chris"
		assertEquals(got , want , t)	

	})

}

func assertEquals( got string , want string , t *testing.T){
	if( got != want){
		t.Errorf(" got %q want %q" , got , want)
	}

}