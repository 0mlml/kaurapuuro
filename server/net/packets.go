package net

import (
	"bytes"
	"encoding/binary"
	"kaurapuuro/sim"
)

var Packets = map[uint8]PacketDef{
	0: &UnknownPacket{},
	1: &UserCommandPacket{},
	2: &UserCommandsPacket{},
	3: &MapDataPacket{},
	4: &GenerateNewLobbyPacket{},
	5: &GenerateNewLobbyResultPacket{},
	6: &EntityListPacket{},
}

type PacketDef interface {
	Serialize() []byte
	Deserialize([]byte) (interface{}, error)
}

type UnknownPacket struct {
}

func (p *UnknownPacket) Serialize() []byte {
	return nil
}

func (p *UnknownPacket) Deserialize(data []byte) (interface{}, error) {
	return nil, nil
}

// Server to client exclusively
type MapDataPacket struct {
	JSONMapData string
}

func (p *MapDataPacket) Serialize() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, []byte(p.JSONMapData))
	return buf.Bytes()
}

func (p *MapDataPacket) Deserialize(data []byte) (interface{}, error) {
	return nil, nil
}

// Client to server exclusively
type UserCommandPacket struct {
	UserCommand sim.UserCommand
}

func (p *UserCommandPacket) Serialize() []byte {
	return nil
}

func (p *UserCommandPacket) Deserialize(data []byte) (interface{}, error) {
	buf := bytes.NewBuffer(data)
	userCommandPacket := &UserCommandPacket{}
	err := binary.Read(buf, binary.BigEndian, &userCommandPacket.UserCommand.EID)
	if err != nil {
		return userCommandPacket, err
	}
	err = binary.Read(buf, binary.BigEndian, &userCommandPacket.UserCommand.Commands)
	if err != nil {
		return userCommandPacket, err
	}
	err = binary.Read(buf, binary.BigEndian, &userCommandPacket.UserCommand.BarrelYaw)
	if err != nil {
		return userCommandPacket, err
	}
	return userCommandPacket, nil
}

// Server to client exclusively
type UserCommandsPacket struct {
	CommandsCount uint8
	UserCommands  []sim.UserCommand
}

func (p *UserCommandsPacket) Serialize() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, p.CommandsCount)
	for _, userCommand := range p.UserCommands {
		binary.Write(buf, binary.BigEndian, userCommand.EID)
		binary.Write(buf, binary.BigEndian, userCommand.Commands)
		binary.Write(buf, binary.BigEndian, userCommand.BarrelYaw)
	}
	return buf.Bytes()
}

func (p *UserCommandsPacket) Deserialize(data []byte) (interface{}, error) {
	return nil, nil
}

// Client to server exclusively
type GenerateNewLobbyPacket struct {
	Name string
}

func (p *GenerateNewLobbyPacket) Serialize() []byte {
	return nil
}

func (p *GenerateNewLobbyPacket) Deserialize(data []byte) (interface{}, error) {
	buf := bytes.NewBuffer(data)
	generateNewLobbyPacket := &GenerateNewLobbyPacket{}
	err := binary.Read(buf, binary.BigEndian, &generateNewLobbyPacket.Name)
	if err != nil {
		return generateNewLobbyPacket, err
	}
	return generateNewLobbyPacket, nil
}

// Server to client
type GenerateNewLobbyResultPacket struct {
	UUID string
}

type EntityListPacket struct {
	Entities []sim.Entity
}
