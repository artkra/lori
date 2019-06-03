package main

import (
	"lori/lserver"
)

func main() {

	server := lserver.NewLServer(10, 1024)

	server.ListenAndServe()
}
