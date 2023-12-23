package mocking

import (
	"testing"
	"bytes"
	"time"
)
func TestCountdown( t *testing.T){
	buffer := &bytes.Buffer{}

	sleeper := SpySleeper{}

	CountDown(buffer , &sleeper)
	got := buffer.String() 
	want := "3\n2\n1\nGo!"
	if got != want {
		t.Errorf(" got %s  want %s" , got , want)
	}

	if( sleeper.Call != 3){
		t.Errorf(" sleep times got %d , want 3" , sleeper.Call)
	}
}

func TestSpytime( t *testing.T){

	sleepTime := 5 * time.Second

	spyTime := &SpyTime {}
	sleeper := ConfigurableSleeper{sleepTime , spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}

}