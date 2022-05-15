package rover_test

import (
	"reflect"
	"testing"

	"github.com/mzwallow/mars-exploration/internal/app/mission"
	"github.com/mzwallow/mars-exploration/internal/app/rover"
)

type result struct {
	direction direction
	position  position
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

var tests = []struct {
	mapSize uint64
	actions []mission.Action

	expectations []result
}{
	{24, []mission.Action{"R", "F", "L", "F", "L", "L", "F", "R"}, []result{
		{"E", position{0, 0}},
		{"E", position{1, 0}},
		{"N", position{1, 0}},
		{"N", position{1, 1}},
		{"W", position{1, 1}},
		{"S", position{1, 1}},
		{"S", position{1, 0}},
		{"W", position{1, 0}},
	}},
	{5, []mission.Action{"L", "L", "L", "F", "L", "R", "F"}, []result{
		{"W", position{0, 0}},
		{"S", position{0, 0}},
		{"E", position{0, 0}},
		{"E", position{1, 0}},
		{"N", position{1, 0}},
		{"E", position{1, 0}},
		{"E", position{2, 0}},
	}},
	{2, []mission.Action{"R", "F", "F", "L", "L"}, []result{
		{"E", position{0, 0}},
		{"E", position{1, 0}},
		{"E", position{1, 0}},
		{"N", position{1, 0}},
		{"W", position{1, 0}},
	}},
}

func TestMove(t *testing.T) {
	for i, test := range tests {
		rover := rover.NewRover()
		var results []result
		for _, a := range test.actions {
			d, p := rover.Move(a, test.mapSize)

			results = append(results, result{direction(d), position(p)})
		}

		if !reflect.DeepEqual(test.expectations, results) {
			t.Errorf("test: %d, want: <%v> \nbut got: <%v>\n",
				i, test.expectations, results)
		}
	}
}
