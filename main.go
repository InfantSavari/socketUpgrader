package main

import (
	"fmt"

	ss "github.com/SocketUpgrader/signalingserver"
	"github.com/google/uuid"
)

func main() {
	hub := ss.Hub{
		Rooms:              make(map[string]uuid.UUID),
		Broadcaster:        make(map[uuid.UUID]ss.IpInfo),
		Observors:          make(map[string][]string),
		ParticipantAddress: make(map[string]ss.IpInfo),
	}

	fmt.Print(hub.Rooms)
}
