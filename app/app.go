package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type Server struct {
	Addr string
}

func (srv Server) ListenAndServe() error {
	addr := srv.Addr

	if addr == "" {
		addr = ":8000"
	}
	log.Printf("Starting server on %v\n", addr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("--- error accepting connection %v\n", err)
			continue
		}
		log.Printf("+++ accepted connection from %v\n", conn.RemoteAddr())
		go handle(conn)
	}
}

func handle(conn net.Conn) error {
	defer func() {
		log.Printf("~ closing connection from %v\n", conn.RemoteAddr())
		conn.Close()
	}()

	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)

	scanr := bufio.NewScanner(r)

	for {
		scanned := scanr.Scan()

		if !scanned {
			if err := scanr.Err(); err != nil {
				log.Printf("%v(%v)", err, conn.RemoteAddr())
				return err
			}
			break
		}

		w.WriteString(strings.ToUpper(scanr.Text()) + "\n")
		w.Flush()
	}
	return nil
}

func main() {
	server := Server{}
	server.ListenAndServe()
}
