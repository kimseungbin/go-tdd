package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	durationA := measureResponseTimer(a)
	durationB := measureResponseTimer(b)

	if durationA < durationB {
		return a
	}

	return b
}

func measureResponseTimer(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
