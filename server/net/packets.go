package net

import (
	"bytes"
	"encoding/binary"
)

var Packets = map[uint8]PacketDef{
	0:  &UnknownPacket{},
	81: &DrawNewPointPacket{},
}

type PacketDef interface {
	Serialize() []byte
	Deserialize([]byte) interface{}
}

type UnknownPacket struct {
}

func (p *UnknownPacket) Serialize() []byte {
	return nil
}

func (p *UnknownPacket) Deserialize(data []byte) interface{} {
	return nil
}

type DrawNewPointPacket struct {
	X     uint16
	Y     uint16
	Color uint8
}

func (p *DrawNewPointPacket) Serialize() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, p.X)
	binary.Write(buf, binary.BigEndian, p.Y)
	binary.Write(buf, binary.BigEndian, p.Color)
	return buf.Bytes()
}

func (p *DrawNewPointPacket) Deserialize(data []byte) interface{} {
	buf := bytes.NewBuffer(data)
	drawPoint := &DrawNewPointPacket{}
	binary.Read(buf, binary.BigEndian, &drawPoint.X)
	binary.Read(buf, binary.BigEndian, &drawPoint.Y)
	binary.Read(buf, binary.BigEndian, &drawPoint.Color)
	return drawPoint
}
