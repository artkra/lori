package lserver

// Dispatcher is a map, carrying info about routing (sender-receiver)
type Dispatcher struct {
	GuestBook *map[string]string
	Router    *map[string]*Conn
}
