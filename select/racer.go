package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTume(a)
	bDuration := measureResponseTume(b)

	if aDuration > bDuration {
		return a
	}

	return b
}

func measureResponseTume(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
