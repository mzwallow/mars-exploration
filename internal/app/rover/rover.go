package rover

import "github.com/mzwallow/mars-exploration/internal/app/mission"

type Rover struct {
	Pos position
	Dir direction
}

type position struct {
	X uint64
	Y uint64
}

type direction string

const (
	North direction = "N"
	East  direction = "E"
	West  direction = "W"
	South direction = "S"
)

func NewRover() *Rover {
	return &Rover{
		Pos: position{X: 0, Y: 0},
		Dir: North,
	}
}

func (r *Rover) Move(action mission.Action, mapSize uint64) (direction, position) {
	switch action {
	case mission.Forward:
		switch r.Dir {
		case North:
			if r.Pos.Y == mapSize-1 {
				break
			}
			r.Pos.Y++
		case South:
			if r.Pos.Y == 0 {
				break
			}
			r.Pos.Y--
		case East:
			if r.Pos.X == mapSize-1 {
				break
			}
			r.Pos.X++
		case West:
			if r.Pos.X == 0 {
				break
			}
			r.Pos.X--
		}
	case mission.Right:
		switch r.Dir {
		case North:
			r.Dir = East
		case South:
			r.Dir = West
		case East:
			r.Dir = South
		case West:
			r.Dir = North
		}
	case mission.Left:
		switch r.Dir {
		case North:
			r.Dir = West
		case South:
			r.Dir = East
		case East:
			r.Dir = North
		case West:
			r.Dir = South
		}
	}

	return r.Dir, r.Pos
}
