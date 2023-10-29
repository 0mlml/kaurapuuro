package sim

import (
	"encoding/json"
	"strconv"
)

// Map positions are stored as relative to map width/height: e.g. w*0.5
func evaluateMapDimension(dimension string) (result float64, ok bool) {
	base := 0.0
	switch dimension[0] {
	case 'w':
		base = WIDTH
	case 'h':
		base = HEIGHT
	default:
		return 0.0, false
	}

	scale, err := strconv.ParseFloat(dimension[2:], 64)
	if err != nil {
		return 0.0, false
	}

	switch dimension[1] {
	case '*':
		return base * scale, true
	case '/':
		if scale == 0 {
			return 0.0, false
		}
		return base / scale, true
	default:
		return 0.0, false
	}
}

type MapFormatLine struct {
	X1    string `json:"x1"`
	Y1    string `json:"y1"`
	X2    string `json:"x2"`
	Y2    string `json:"y2"`
	Color string `json:"color"`
}

func (l *MapFormatLine) ToLine() Line {
	x1, ok := evaluateMapDimension(l.X1)
	if !ok {
		return Line{}
	}

	y1, ok := evaluateMapDimension(l.Y1)
	if !ok {
		return Line{}
	}

	x2, ok := evaluateMapDimension(l.X2)
	if !ok {
		return Line{}
	}

	y2, ok := evaluateMapDimension(l.Y2)
	if !ok {
		return Line{}
	}

	return Line{
		Start: Vector{
			X: x1,
			Y: y1,
		},
		End: Vector{
			X: x2,
			Y: y2,
		},
	}
}

type MapFormatBox struct {
	X      string `json:"x"`
	Y      string `json:"y"`
	Width  string `json:"width"`
	Height string `json:"height"`
	Color  string `json:"color"`
}

func (b *MapFormatBox) ToLines() []Line {
	x, ok := evaluateMapDimension(b.X)
	if !ok {
		return nil
	}

	y, ok := evaluateMapDimension(b.Y)
	if !ok {
		return nil
	}

	width, ok := evaluateMapDimension(b.Width)
	if !ok {
		return nil
	}

	height, ok := evaluateMapDimension(b.Height)
	if !ok {
		return nil
	}

	return []Line{
		{
			Start: Vector{
				X: x,
				Y: y,
			},
			End: Vector{
				X: x + width,
				Y: y,
			},
		},
		{
			Start: Vector{
				X: x + width,
				Y: y,
			},
			End: Vector{
				X: x + width,
				Y: y + height,
			},
		},
		{
			Start: Vector{
				X: x + width,
				Y: y + height,
			},
			End: Vector{
				X: x,
				Y: y + height,
			},
		},
		{
			Start: Vector{
				X: x,
				Y: y + height,
			},
			End: Vector{
				X: x,
				Y: y,
			},
		},
	}
}

type MapFormatSpawnPoint struct {
	X string `json:"x"`
	Y string `json:"y"`

	ID    uint8 `json:"id"`
	Flags uint8 `json:"flags"`
}

func (sp *MapFormatSpawnPoint) ToSpawnPoint() SpawnPoint {
	x, ok := evaluateMapDimension(sp.X)
	if !ok {
		return SpawnPoint{}
	}

	y, ok := evaluateMapDimension(sp.Y)
	if !ok {
		return SpawnPoint{}
	}

	return SpawnPoint{
		Vector{
			X: x,
			Y: y,
		},
		sp.ID,
		sp.Flags,
	}
}

type MapFormat struct {
	Lines       []MapFormatLine       `json:"lines"`
	Boxes       []MapFormatBox        `json:"boxes"`
	SpawnPoints []MapFormatSpawnPoint `json:"spawnPoints"`
}

func (m *Map) ParseMapFromJson(jsonMap string) {
	mapFormat := MapFormat{}
	err := json.Unmarshal([]byte(jsonMap), &mapFormat)
	if err != nil {
		panic(err)
	}

	for _, line := range mapFormat.Lines {
		m.Lines = append(m.Lines, line.ToLine())
	}

	for _, box := range mapFormat.Boxes {
		m.Lines = append(m.Lines, box.ToLines()...)
	}

	for _, spawnPoint := range mapFormat.SpawnPoints {
		m.SpawnPoints = append(m.SpawnPoints, spawnPoint.ToSpawnPoint())
	}
}
