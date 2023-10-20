package main

import (
	"kaurapuuro/net"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	net.RegisterPacketHandler(81, func(packet *net.Packet, conn *websocket.Conn) {
		drawPoint := packet.Data.(*net.DrawNewPointPacket)

		log.Printf("Received new point: %d, %d, %d", drawPoint.X, drawPoint.Y, drawPoint.Color)

		net.BroadcastPacket(81, drawPoint.Serialize())
	})

	log.Println("Starting server...")
	net.StartServer(":8080")
}
