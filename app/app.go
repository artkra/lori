package main

import (
	"lori/lserver"
	"time"
)

func main() {
	server := lserver.Server{
		IdleTimeout:   10 * time.Second,
		MaxReadBuffer: 1024,
		Dispatch:      make(map[int]int, 1000),
	}

	server.ListenAndServe()
}
