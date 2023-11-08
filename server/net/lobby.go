package net

import (
	"kaurapuuro/sim"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	LOBBY_MAX_EMPTY_TIME = 10000
)

type Lobby struct {
	UUID    string
	Name    string
	MapJSON string

	Game *sim.Game

	Clients        []*Client
	timeFirstEmpty int64
}

var lobbies = make(map[string]*Lobby)

func ClientExistsInLobby(client *Client) (ok bool, l *Lobby) {
	for _, lobby := range lobbies {
		for _, c := range lobby.Clients {
			if c.UUID == client.UUID {
				return true, lobby
			}
		}
	}

	return false, nil
}

func NewLobby(name string) *Lobby {
	lobby := &Lobby{
		UUID: generateUUID(),
		Name: name,
	}

	lobbies[lobby.UUID] = lobby

	return lobby
}

func (l *Lobby) AddClient(client *Client) {
	l.Clients = append(l.Clients, client)
}

func (l *Lobby) RemoveClient(client *Client) {
	for i, c := range l.Clients {
		if c.UUID == client.UUID {
			l.Clients = append(l.Clients[:i], l.Clients[i+1:]...)
			return
		}
	}

	if len(l.Clients) == 0 {
		l.timeFirstEmpty = time.Now().UnixMilli()
	}
}

func GetLobbyByUUID(uuid string) (ok bool, l *Lobby) {
	for _, lobby := range lobbies {
		if lobby.UUID == uuid {
			return true, lobby
		}
	}
	return false, nil
}

func GenerateLobbyPacketListener(p *Packet, c *websocket.Conn) {
	if (p.Error != nil) || (p.Data == nil) {
		log.Printf("Error deserializing packet: %s", p.Error)
	}

	resultPacket := &GenerateNewLobbyResultPacket{}

	client := GetClientByConn(c)

	if client == nil {
		log.Printf("Received packet from unknown client")
		return
	}

	if ok, lobby := ClientExistsInLobby(client); ok {
		log.Printf("Client %s is already in a lobby when creating new lobby. Dropping them from lobby %s", client.UUID, lobby.UUID)
		lobby.RemoveClient(client)
	}

	packet := p.Data.(*GenerateNewLobbyPacket)

	if packet.Name == "" {
		log.Printf("Received empty lobby name")
		client.SendPacket(5, resultPacket.Serialize())
		return
	}

	lobby := NewLobby(packet.Name)

	log.Printf("Generated new lobby with UUID: %s", lobby.UUID)

	resultPacket.UUID = lobby.UUID
	client.SendPacket(5, resultPacket.Serialize())
}
