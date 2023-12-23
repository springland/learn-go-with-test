package mocking
import (
	"fmt"
	"io"
	"os"
	"time"
)
func main() {

	defaultSleeper := DefaultSleeper{}

	CountDown(os.Stdout , &defaultSleeper)
}

type Sleeper interface {

	Sleep()
}

type SpySleeper struct {
	Call int
}

func (sleeper *SpySleeper) Sleep() {
	sleeper.Call++
}

type DefaultSleeper struct {

}


func (sleeper *DefaultSleeper)  Sleep() {
	time.Sleep(1 * time.Second)
}



func CountDown( w io.Writer , sleeper  Sleeper){

	for index:= 3 ; index >0 ; index-- {
		fmt.Fprintf(w , "%d\n" , index)
		sleeper.Sleep()
	}

	fmt.Fprintf(w , "Go!")
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

type SpyTime struct {

	durationSlept time.Duration
}

func ( spyTime *SpyTime) Sleep(duration time.Duration){
	spyTime.durationSlept = duration
}

func (sleeper *ConfigurableSleeper) Sleep(){
	sleeper.sleep(sleeper.duration)
}