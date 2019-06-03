package lserver

// Dispatcher is a map, carrying info about routing (sender-receiver)
type Dispatcher struct {
	Router *map[string]*Conn
}
