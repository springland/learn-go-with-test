package websiteracer

import (
	"time"
	"net/http"

)
func SequentialRacer( a , b string) (winner string){

	aStart :=  time.Now()
	http.Get(a)
	aDuration := time.Since(aStart)


	bStart := time.Now()
	http.Get(b)

	bDuration := time.Since(bStart)

	if( aDuration < bDuration){
		return a
	}

	return b	
}

func ConcurrentRacer( a , b string)(winner string){

	select{
		
	case  <- ping(a):
		return a

	case <- ping(b):
		return b	
	}

}


func ping( url string) chan  struct {} {

	ch := make(chan struct{})

	go func () {
		http.Get(url)
		close(ch)
	}()

	return ch
}
