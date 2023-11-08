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
	UserCommand sim.NetUserCommand
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
	UserCommands  []sim.NetUserCommand
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
	UUID string // UUID of the lobby, empty if failed
}

func (p *GenerateNewLobbyResultPacket) Serialize() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, p.UUID)
	return buf.Bytes()
}

func (p *GenerateNewLobbyResultPacket) Deserialize(data []byte) (interface{}, error) {
	return nil, nil
}

// Server to client
type EntityListPacket struct {
	Entities sim.NetEntityList
}

func (p *EntityListPacket) Serialize() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, p.Entities.TankCount)
	for _, tank := range p.Entities.Tanks {
		binary.Write(buf, binary.BigEndian, tank.EID)
		binary.Write(buf, binary.BigEndian, tank.Position.X)
		binary.Write(buf, binary.BigEndian, tank.Position.Y)
		binary.Write(buf, binary.BigEndian, tank.Velocity.X)
		binary.Write(buf, binary.BigEndian, tank.Velocity.Y)
		binary.Write(buf, binary.BigEndian, tank.EFlags)
		binary.Write(buf, binary.BigEndian, tank.BarrelYaw)
		binary.Write(buf, binary.BigEndian, tank.Flags)
	}
	binary.Write(buf, binary.BigEndian, p.Entities.ProjectileCount)
	for _, projectile := range p.Entities.Projectiles {
		binary.Write(buf, binary.BigEndian, projectile.EID)
		binary.Write(buf, binary.BigEndian, projectile.Position.X)
		binary.Write(buf, binary.BigEndian, projectile.Position.Y)
		binary.Write(buf, binary.BigEndian, projectile.Velocity.X)
		binary.Write(buf, binary.BigEndian, projectile.Velocity.Y)
		binary.Write(buf, binary.BigEndian, projectile.EFlags)
		binary.Write(buf, binary.BigEndian, projectile.BaseVelocity.X)
		binary.Write(buf, binary.BigEndian, projectile.BaseVelocity.Y)
		binary.Write(buf, binary.BigEndian, projectile.Bounces)
		binary.Write(buf, binary.BigEndian, projectile.OwnerEID)
	}
	binary.Write(buf, binary.BigEndian, p.Entities.MineCount)
	for _, mine := range p.Entities.Mines {
		binary.Write(buf, binary.BigEndian, mine.EID)
		binary.Write(buf, binary.BigEndian, mine.Position.X)
		binary.Write(buf, binary.BigEndian, mine.Position.Y)
		binary.Write(buf, binary.BigEndian, mine.Velocity.X)
		binary.Write(buf, binary.BigEndian, mine.Velocity.Y)
		binary.Write(buf, binary.BigEndian, mine.EFlags)
		binary.Write(buf, binary.BigEndian, mine.ArmTime)
		binary.Write(buf, binary.BigEndian, mine.OwnerEID)
	}
	return buf.Bytes()
}

func (p *EntityListPacket) Deserialize(data []byte) (interface{}, error) {
	return nil, nil
}
