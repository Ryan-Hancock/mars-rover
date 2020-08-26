package main

import (
	"fmt"
	"mars-rover/rover"
	"strconv"
)

// Start takes the command lines and inputs them into the rover Controllers.
func Start(lines [][]string) {
	if len(lines) < 3 {
		return
	}

	// first line to build the plateau, example: (5, 5)
	var plataue = lines[0]
	x := convertToInt(plataue[0])
	y := convertToInt(plataue[1])

	squad := rover.SetupMission(x, y)

	remainingLines := lines[1:]
	// check if the rest of the array are rover position and command lines
	if len(remainingLines)%2 != 0 {
		return
	}

	// second and third line, example: (1 2 N) (LMLMLMLMM)...
	for i := 0; i < len(remainingLines); i += 2 {
		con, err := makeRoverAndMove(squad, remainingLines[i], remainingLines[i+1])
		if err != nil {
			continue
		}
		fmt.Println(con)
	}
}

func makeRoverAndMove(s *rover.Squad, r []string, commands []string) (*rover.Controller, error) {
	x := convertToInt(r[0])
	y := convertToInt(r[1])
	dir := rover.StrToDirection(r[2])

	controller, err := s.AddRover(rover.Grid{x}, rover.Grid{y}, dir)
	if err != nil {
		return nil, err
	}

	for _, command := range commands[0] {
		controller.Operate(rover.Command(command))
	}

	return controller, nil
}

func convertToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return i
}
