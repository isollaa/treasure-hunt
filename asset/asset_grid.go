package asset

import (
	"fmt"
)

func buildGrid(coordinateChar Coordinate) (found bool) {
	level := totalTreasure - len(coordinateTreasures)
	fmt.Printf("Treasure - Hunt [%d / %d]\n", level, totalTreasure)
	fmt.Println("")
	for i := 0; i < yGridLength; i++ {
		for j := 0; j < xGridLength; j++ {
			if isEqualCharacterCoordinate(j, i, coordinateChar) && isEqualTreasure(j, i) {
				fmt.Print("X")
				found = true
			} else if isEqualCharacterCoordinate(j, i, coordinateChar) {
				fmt.Print("X")
			} else if _, ok := coordinateObstacles[Coordinate{j, i}]; ok {
				fmt.Print("#")
			} else if isEqualTreasure(j, i) {
				fmt.Print("$")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println("")
	}

	return found
}
