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
