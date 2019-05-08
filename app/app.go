package main

import (
	"bufio"
	"io"
	"log"
	"lori/lserver"
	"net"
	"strings"
	"time"
)

type Server struct {
	Addr          string
	IdleTimeout   time.Duration
	MaxReadBuffer int64
	Dispatch      map[int]int
}

type Conn struct {
	net.Conn
	IdleTimeout   time.Duration
	MaxReadBuffer int64
	Dispatch      *map[int]int
}

func (c *Conn) Write(p []byte) (int, error) {
	c.UpdateDeadline()
	return c.Conn.Write(p)
}

func (c *Conn) Read(b []byte) (int, error) {
	c.UpdateDeadline()
	r := io.LimitReader(c.Conn, c.MaxReadBuffer)
	log.Println((*c.Dispatch)[1])
	return r.Read(b)
}

func (c *Conn) UpdateDeadline() {
	idleDeadline := time.Now().Add(c.IdleTimeout)
	c.Conn.SetDeadline(idleDeadline)
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
		newConn, err := listener.Accept()
		if err != nil {
			log.Printf("--- error accepting connection %v\n", err)
			continue
		}
		conn := &Conn{
			Conn:          newConn,
			IdleTimeout:   srv.IdleTimeout,
			MaxReadBuffer: srv.MaxReadBuffer,
			Dispatch:      &srv.Dispatch,
		}
		conn.SetDeadline(time.Now().Add(conn.IdleTimeout))
		log.Printf("+++ accepted connection from %v\n", conn.RemoteAddr())
		go handle(conn)
	}
}

func handle(conn *Conn) error {
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
	server := Server{
		IdleTimeout:   10 * time.Second,
		MaxReadBuffer: 1024,
		Dispatch:      make(map[int]int, 1000),
	}

	lserver.Shout()
	lserver.QShout()

	server.ListenAndServe()
}
