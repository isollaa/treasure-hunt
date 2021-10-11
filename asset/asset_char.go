package asset

import (
	"fmt"
	"math"
	"time"
	"treasure-hunt/lib"
)

var coordinateCharDefault = Coordinate{1, 4}
var historyStep = []Coordinate{coordinateCharDefault}

func SetCharacterCoordinate(xChar int, yChar int) (found bool) {
	//clear cli to make it like animating
	lib.CallClear()
	coordinateChar := Coordinate{xChar, yChar}
	found = buildGrid(coordinateChar)
	//reset coordinate history if character is on treasure coordinate
	if found {
		historyStep = []Coordinate{coordinateCharDefault}
		fmt.Println("\ntreasure found at", coordinateChar)
	}
	time.Sleep(350 * time.Millisecond)

	return
}

func GetCoordinateCharDefault() Coordinate {
	return coordinateCharDefault
}

func GetHistoryStep() []Coordinate {
	return historyStep
}

func MoveUpCharacter(char *Coordinate) (isFound bool) {
	char.Y--
	return move(char, "U")
}

func MoveDownCharacter(char *Coordinate) (isFound bool) {
	char.Y++
	return move(char, "D")
}

func MoveRightCharacter(char *Coordinate) (isFound bool) {
	char.X++
	return move(char, "R")
}

func isRepeatedStep(coor Coordinate) bool {
	for _, v := range historyStep {
		if v == coor {
			return true
		}
	}

	return false
}

//check is character is one step ahead to treasure
func isCharacterNearTreasure(char *Coordinate) bool {
	if coordinateTreasure.X == char.X {
		if math.Abs(float64(coordinateTreasure.Y-char.Y)) == 1 {
			*char = coordinateTreasure
			return true
		}
	} else if coordinateTreasure.Y == char.Y {
		if math.Abs(float64(coordinateTreasure.X-char.X)) == 1 {
			*char = coordinateTreasure
			return true
		}
	}

	return false
}

func move(char *Coordinate, moveType string) (isFound bool) {
	//check if character moved to obstacle coordinate / to stepped coordinates
	if _, ok := coordinateObstacles[Coordinate{char.X, char.Y}]; ok || isRepeatedStep(*char) {
		switch moveType {
		case "U":
			char.Y++
		case "D":
			char.Y--
		case "R":
			char.X--
		}
		return
	}

	//store success history step
	historyStep = append(historyStep, *char)
	isFound = SetCharacterCoordinate(char.X, char.Y)
	//if character is one step ahead to treasure, auto move the character
	if isCharacterNearTreasure(char) {
		isFound = SetCharacterCoordinate(char.X, char.Y)
	}

	return
}

//get all obstacle coordinates
func isEqualCharacterCoordinate(x int, y int, coordinateChar Coordinate) bool {
	return x == coordinateChar.X && y == coordinateChar.Y
}
