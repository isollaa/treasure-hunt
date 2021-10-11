package core

import "treasure-hunt/asset"

func initAction() {
	//get init position coordinate
	char := asset.GetCoordinateCharDefault()
	//set random treasure position based on clear path
	asset.SetRandomCoordinateTreasure()
	//set character with the grid
	asset.SetCharacterCoordinate(char.X, char.Y)
	//move character to treasure coordinate
	moveCharacter(char)
}

func moveCharacter(char asset.Coordinate) {
	// temp variable to store last movement with number of loop
	tempCoordinate := map[int]asset.Coordinate{}
	counter := 0

	for {
		counter++
		if char.Y > asset.GetCoordinateTreasure().Y {
			if isFound := asset.MoveUpCharacter(&char); isFound {
				return
			}
		}

		// tebreak loop if character is on the same position
		if isCharacterIdle(tempCoordinate, counter) {
			break
		}
	}

	//reset counter
	counter = 0
	for {
		counter++
		if isFound := asset.MoveRightCharacter(&char); isFound {
			break
		}

		if isFound := asset.MoveDownCharacter(&char); isFound {
			break
		}

		if isCharacterIdle(tempCoordinate, counter) {
			if isFound := asset.MoveUpCharacter(&char); isFound {
				break
			}
		}
	}
}

func isCharacterIdle(tempCoordinate map[int]asset.Coordinate, idx int) (idle bool) {
	historyStep := asset.GetHistoryStep()
	tempCoordinate[idx] = historyStep[len(historyStep)-1]

	if len(tempCoordinate) > 1 {
		if tempCoordinate[idx] == tempCoordinate[idx-1] {
			return true
		}
	}

	return
}
