package sim

type Vector struct {
	X, Y float32
}

type TankType struct {
	DisplayName string
	ID          string

	BodySpeed   float32
	BarrelSpeed float32

	BulletType ProjectileType
}

type Player struct {
	Position  Vector
	Velocity  Vector
	Yaw       float32
	BarrelYaw float32

	Buttons ButtonState

	IsBot bool

	Type        TankType
	BodySpeed   float32
	BarrelSpeed float32
}

type ProjectileType int

const (
	BulletProjectileType ProjectileType = iota
	RocketProjectileType
	MineProjectileType
)

type Projectile struct {
	Type    ProjectileType
	ArmTime int
	Yaw     float32
	// Other specific projectile fields...
}

type ButtonState struct {
	Forward bool
	Back    bool
	Left    bool
	Right   bool
	Fire    bool
	AltFire bool
}

type MapSpawnpoint struct {
	Position Vector  `json:"position"`
	Yaw      float32 `json:"yaw"`
	Team     int     `json:"team"`
	TankType string  `json:"tankType"`

	MaxUses    int `json:"maxUses"`
	MaxHp      int `json:"maxHp"`
	CooldownMs int `json:"cooldownMs"`
}

type MapLine struct {
	Start Vector `json:"start"`
	End   Vector `json:"end"`

	Color string `json:"color"`
}

type MapBox struct {
	Position Vector  `json:"position"`
	Width    float32 `json:"width"`
	Height   float32 `json:"depth"`
	Hollow   bool    `json:"hollow"`

	Color string `json:"color"`
}

type MapEnemySpawnpoint struct {
	Position Vector   `json:"position"`
	Yaw      float32  `json:"yaw"`
	Team     int      `json:"team"`
	TankType TankType `json:"tankType"`
}

type Map struct {
	Spawnpoints []MapSpawnpoint
}

type MatchSpawnpoint struct {
	MapSpawnpoint

	Hp       int `json:"hp"`
	Uses     int `json:"uses"`
	Cooldown int `json:"cooldown"`
}

type Match struct {
}
