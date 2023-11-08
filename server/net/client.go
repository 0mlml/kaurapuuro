package net

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	UUID     string
	Nickname string
	Conn     *websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		UUID: generateUUID(),
		Conn: conn,
	}
}

func GetClientByConn(conn *websocket.Conn) *Client {
	for _, client := range clients {
		if client.Conn == conn {
			return client
		}
	}

	return nil
}

func (c *Client) SendPacket(packetType uint8, data []byte) error {
	payload := append([]byte{packetType}, data...)
	return c.Conn.WriteMessage(websocket.BinaryMessage, payload)
}
