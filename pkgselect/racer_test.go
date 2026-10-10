package pkgselect

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare speed of servers", func(t *testing.T) {
		slowServer := makeDelayServer(8 * time.Millisecond)
		fastServer := makeDelayServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns an eror if a server doesnt respond within 10s", func(t *testing.T) {
		serverA := makeDelayServer(11 * time.Millisecond)
		serverB := makeDelayServer(12 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)
		if err == nil {
			t.Errorf("expected an er but didnt got one")
		}
	})
}

func makeDelayServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}
