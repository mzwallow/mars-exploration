package operation

import (
	"fmt"

	"github.com/mzwallow/mars-exploration/internal/app/mission"
	"github.com/mzwallow/mars-exploration/internal/app/rover"
)

type Operation struct {
	mission *mission.Mission
	rover   *rover.Rover
}

func NewOperation(mission *mission.Mission, rover *rover.Rover) *Operation {
	return &Operation{mission: mission, rover: rover}
}

func (o *Operation) Run() {
	fmt.Printf("%d\t%s:%d,%d\n",
		o.mission.MapSize,
		o.rover.Dir,
		o.rover.Pos.X,
		o.rover.Pos.Y)

	for _, action := range o.mission.Actions {
		dir, pos := o.rover.Move(action, o.mission.MapSize)

		fmt.Printf("%v\t%v:%d,%d\n",
			action,
			dir,
			pos.X,
			pos.Y)
	}
}
