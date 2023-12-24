package websiteracer

import (
	"testing"
	"time"
	"net/http"
	"net/http/httptest"

)
func TestWebsiteRacer(t *testing.T){

	slowURL := "http://facebook.com"
	fastURL := "http://www.quii.dev"

	want := fastURL

	got := SequentialRacer(slowURL , fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	} 
}


func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestConcurrentRacer( t *testing.T){

	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := ConcurrentRacer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}