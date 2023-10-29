package net

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Packet struct {
	Type  uint8
	Data  interface{}
	Error error
}

type PacketHandler func(*Packet, *websocket.Conn)

var packetHandlers = make(map[uint8]PacketHandler)

func RegisterPacketHandler(packetType uint8, handler PacketHandler) {
	packetHandlers[packetType] = handler
}

func UnregisterPacketHandler(packetType uint8) {
	delete(packetHandlers, packetType)
}

func DeserializePacket(data []byte) *Packet {
	packetType := data[0]
	packetData := data[1:]

	packet := &Packet{
		Type: packetType,
	}

	if handler, exists := Packets[packetType]; exists {
		packet.Data, packet.Error = handler.Deserialize(packetData)
	} else {
		fmt.Printf("Received unknown packet type %d", packetType)
	}

	return packet
}

var clients = make(map[string]*Client)

func BroadcastPacket(packetType uint8, data []byte) {
	for _, client := range clients {
		client.SendPacket(packetType, data)
	}
}

func handleConnection(conn *websocket.Conn) {
	defer conn.Close()
	client := NewClient(conn)

	fmt.Printf("New client connected with UUID: %s, From: %s\n", client.UUID, conn.RemoteAddr())

	clients[client.UUID] = client
	defer delete(clients, client.UUID)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}

		packet := DeserializePacket(message)

		fmt.Printf("Received packet type %d from client %s\nRaw: %v\n", packet.Type, client.UUID, message)

		if handler, exists := packetHandlers[packet.Type]; exists {
			handler(packet, conn)
		}
	}
}

func StartServer(addr string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		go handleConnection(conn)
	})

	log.Fatal(http.ListenAndServe(addr, nil))
}
