package sim

const (
	TICKRATE = 30
	DELTA    = 1.0 / TICKRATE
	WIDTH    = 2000
	HEIGHT   = 2000

	DEFAULTMAP = ``
)

func getLineIntersection(line1, line2 Line) (vec Vector, ok bool) {
	delta1 := Vector{
		X: line1.End.X - line1.Start.X,
		Y: line1.End.Y - line1.Start.Y,
	}

	delta2 := Vector{
		X: line2.End.X - line2.Start.X,
		Y: line2.End.Y - line2.Start.Y,
	}

	// Check for degenerate lines (points)
	if (delta1.X == 0 && delta1.Y == 0) || (delta2.X == 0 && delta2.Y == 0) {
		return Vector{}, false
	}

	denominator := delta2.Y*delta1.X - delta2.X*delta1.Y

	// Lines are parallel or coincident
	if denominator == 0 {
		return Vector{}, false
	}

	ua := (delta2.X*(line1.Start.Y-line2.Start.Y) - delta2.Y*(line1.Start.X-line2.Start.X)) / denominator
	ub := (delta1.X*(line1.Start.Y-line2.Start.Y) - delta1.Y*(line1.Start.X-line2.Start.X)) / denominator

	// Intersection outside the line segments
	if ua < 0 || ua > 1 || ub < 0 || ub > 1 {
		return Vector{}, false
	}

	return Vector{
		X: line1.Start.X + ua*delta1.X,
		Y: line1.Start.Y + ua*delta1.Y,
	}, true
}

func (g *Game) tick() {

}
