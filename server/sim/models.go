package sim

type Vector struct {
	X, Y float64
}

type Line struct {
	Start Vector
	End   Vector
}

type UserCommand struct {
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

func (uc *UserCommand) SetForward() {
	uc.Commands |= ForwardBit
}

func (uc *UserCommand) SetBack() {
	uc.Commands |= BackBit
}

func (uc *UserCommand) SetLeft() {
	uc.Commands |= LeftBit
}

func (uc *UserCommand) SetRight() {
	uc.Commands |= RightBit
}

func (uc *UserCommand) SetFire() {
	uc.Commands |= FireBit
}

func (uc *UserCommand) SetAltFire() {
	uc.Commands |= AltFireBit
}

func (uc *UserCommand) IsForward() bool {
	return uc.Commands&ForwardBit != 0
}

func (uc *UserCommand) IsBack() bool {
	return uc.Commands&BackBit != 0
}

func (uc *UserCommand) IsLeft() bool {
	return uc.Commands&LeftBit != 0
}

func (uc *UserCommand) IsRight() bool {
	return uc.Commands&RightBit != 0
}

func (uc *UserCommand) IsFire() bool {
	return uc.Commands&FireBit != 0
}

func (uc *UserCommand) IsAltFire() bool {
	return uc.Commands&AltFireBit != 0
}

const (
	FlagAlive = 1 << iota
	FlagInvincible
)

type Entity struct {
	EID      uint16
	Position Vector
	Velocity Vector
	Flags    uint8
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

type Tank struct {
	Entity
	BarrelYaw uint16 // Tenths of degree e.g. 37.1 -> 371
	Flags     uint8  // Team and Type
}

type Bullet struct {
	Entity
	BaseVelocity Vector
	Bounces      uint8
	OwnerEID     uint16
}

type Mine struct {
	Entity
	ArmTime  uint16
	OwnerEID uint16
}

type EntityList struct {
	TankCount       uint8
	Tanks           []Tank
	ProjectileCount uint8
	Projectiles     []Bullet
	MineCount       uint8
	Mines           []Mine
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
	Entities EntityList
	Map      Map
}
