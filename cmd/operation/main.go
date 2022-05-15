package main

import (
	"fmt"
	"os"

	"github.com/mzwallow/mars-exploration/internal/app/mission"
	"github.com/mzwallow/mars-exploration/internal/app/operation"
	"github.com/mzwallow/mars-exploration/internal/app/rover"
)

func main() {
	fmt.Println(`Instructions:
	1. Put your 'file' in to 'files' directory.
	2. Type your name of that 'file' with extension e.g. 'm1.txt' when promt.
	3. Hit Enter.`)

	var fileName string
	fmt.Printf("\nEnter your file name: ")
	fmt.Scanln(&fileName)

	path := fmt.Sprintf("files/%s", fileName)
	missionFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer missionFile.Close()

	mission := mission.NewMission(missionFile)
	rover := rover.NewRover()
	operation := operation.NewOperation(mission, rover)
	fmt.Printf("\nResult:\n")
	operation.Run()
}
