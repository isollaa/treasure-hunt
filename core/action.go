package core

import "treasure-hunt/asset"

func initAction() {
	char := asset.GetCoordinateCharDefault()
	asset.SetRandomCoordinateTreasure()
	asset.SetCharacterCoordinate(char.X, char.Y)
	moveCharacter(char)
}

func moveCharacter(char asset.Coordinate) {
	tempCoordinate := map[int]asset.Coordinate{}
	counter := 0

	for {
		counter++
		if char.Y > asset.GetCoordinateTreasure().Y {
			if isFound := asset.MoveUpCharacter(&char); isFound {
				return
			}
		}

		if isCharacterIdle(tempCoordinate, counter) {
			break
		}
	}

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
