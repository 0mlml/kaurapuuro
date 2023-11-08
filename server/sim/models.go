package sim

type Vector struct {
	X, Y float64
}

type Line struct {
	Start Vector
	End   Vector
}

type NetUserCommand struct {
	EID       uint16 // Entity ID
	Commands  uint8  // bitfield for Forward, Back, Left, Right, Fire, AltFire
	BarrelYaw uint16 // Tenths of degree e.g. 37.1 -> 371
}

const (
	ForwardBit = 1 << iota
	BackBit
	LeftBit
	RightBit
	FireBit
	AltFireBit
)

func (uc *NetUserCommand) SetForward() {
	uc.Commands |= ForwardBit
}

func (uc *NetUserCommand) SetBack() {
	uc.Commands |= BackBit
}

func (uc *NetUserCommand) SetLeft() {
	uc.Commands |= LeftBit
}

func (uc *NetUserCommand) SetRight() {
	uc.Commands |= RightBit
}

func (uc *NetUserCommand) SetFire() {
	uc.Commands |= FireBit
}

func (uc *NetUserCommand) SetAltFire() {
	uc.Commands |= AltFireBit
}

func (uc *NetUserCommand) IsForward() bool {
	return uc.Commands&ForwardBit != 0
}

func (uc *NetUserCommand) IsBack() bool {
	return uc.Commands&BackBit != 0
}

func (uc *NetUserCommand) IsLeft() bool {
	return uc.Commands&LeftBit != 0
}

func (uc *NetUserCommand) IsRight() bool {
	return uc.Commands&RightBit != 0
}

func (uc *NetUserCommand) IsFire() bool {
	return uc.Commands&FireBit != 0
}

func (uc *NetUserCommand) IsAltFire() bool {
	return uc.Commands&AltFireBit != 0
}

const (
	FlagAlive = 1 << iota
	FlagInvincible
)

type NetEntity struct {
	EID      uint16
	Position Vector
	Velocity Vector
	EFlags   uint8
}

const (
	TankCommon = iota
	TankFast
	TankRocket
	TankBounce
	TankFive
	TankSix
	TankSeven
	TankEight
)

const (
	TeamOne = (iota + 1) << 4
	TeamTwo
	TeamThree
	TeamFour
	TeamFive
	TeamSix
	TeamSeven
	TeamEight
)

type NetTank struct {
	NetEntity
	BarrelYaw uint16 // Tenths of degree e.g. 37.1 -> 371
	Flags     uint8  // Team and Type
}

type NetBullet struct {
	NetEntity
	BaseVelocity Vector
	Bounces      uint8
	OwnerEID     uint16
}

type NetMine struct {
	NetEntity
	ArmTime  uint16
	OwnerEID uint16
}

type NetEntityList struct {
	TankCount       uint8
	Tanks           []NetTank
	ProjectileCount uint8
	Projectiles     []NetBullet
	MineCount       uint8
	Mines           []NetMine
}

type SpawnPoint struct {
	Position Vector
	ID       uint8
	Flags    uint8
}

type Map struct {
	Lines       []Line
	SpawnPoints []SpawnPoint
}

type Game struct {
	Dimensions Vector // Width and height of the game world

	Entities NetEntityList
	Map      Map
}
