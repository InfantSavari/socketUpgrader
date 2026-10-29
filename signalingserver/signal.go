package signalingserver

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

type IpInfo struct {
	Address string
	Port    int
}

type Hub struct {
	Mu                 sync.RWMutex
	Rooms              map[string]uuid.UUID
	Broadcaster        map[uuid.UUID]IpInfo
	Observors          map[string][]string
	ParticipantAddress map[string]IpInfo
}

func (hb *Hub) IsRoomNameAvailable(name string) bool {
	result := hb.Rooms[name]
	if result != uuid.Nil {
		return false
	}
	return true
}

func GenerateUID() uuid.UUID {
	result, err := uuid.NewV7()
	if err != nil {
		return uuid.Nil
	}
	return result
}

func (hb *Hub) RecieveIPAddress(w http.ResponseWriter, r *http.Request) {

}

func (hb *Hub) SendIPToBroadcaster() {

}

func (hb *Hub) CreateRoom(w http.ResponseWriter, r *http.Request) {

	type RoomInfo struct {
		RoomName  string `json:"roomname"`
		IpAddress string `json:"IpAddress"`
		Port      int    `json:"port"`
	}

	var payload RoomInfo
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		log.Println("Data decoding at Room Creation failed....")
		return
	}

	var id uuid.UUID = uuid.Nil
	for id == uuid.Nil {
		id = GenerateUID()
	}

	hb.Mu.Lock()
	if hb.IsRoomNameAvailable(payload.RoomName) {
		hb.Rooms[payload.RoomName] = id
		hb.Broadcaster[id] = IpInfo{
			Address: payload.IpAddress,
			Port:    payload.Port,
		}
	}

}
