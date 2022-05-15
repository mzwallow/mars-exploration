package mission

import (
	"bufio"
	"os"
	"strconv"
)

type Mission struct {
	MapSize uint64
	Actions []Action
}

type Action string

const (
	Forward Action = "F"
	Right   Action = "R"
	Left    Action = "L"
)

func NewMission(missionFile *os.File) *Mission {
	var mapSize uint64
	var actions []Action

	scanner := bufio.NewScanner(missionFile)
	isFirstLine := true
	for scanner.Scan() {
		if isFirstLine {
			isFirstLine = false
			mapSize, _ = strconv.ParseUint(scanner.Text(), 10, 64)
			continue
		}
		actions = append(actions, Action(scanner.Text()))
	}

	return &Mission{MapSize: mapSize, Actions: actions}
}
