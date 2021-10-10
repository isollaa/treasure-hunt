package asset

import (
	"math/rand"
	"time"
)

var coordinateTreasure Coordinate
var coordinateTreasures = []Coordinate{}
var totalTreasure int

func setTreasureCoordinates() {
	for i := 0; i < yGridLength; i++ {
		for j := 0; j < xGridLength; j++ {
			if _, ok := coordinateObstacles[Coordinate{j, i}]; ok || j == coordinateCharDefault.X && i == coordinateCharDefault.Y {
				continue
			} else {
				coordinateTreasures = append(coordinateTreasures, Coordinate{j, i})
			}
		}
	}

	totalTreasure = len(coordinateTreasures)
}

func SetRandomCoordinateTreasure() {
	rand.Seed(time.Now().Unix())
	idxRandom := rand.Intn(len(coordinateTreasures))
	coordinateTreasure = coordinateTreasures[idxRandom]
	coordinateTreasures = removeCoordinateTreasureIndex(coordinateTreasures, idxRandom)
}

func GetCoordinateTreasure() Coordinate {
	return coordinateTreasure
}

func GetCoordinateTreasures() []Coordinate {
	return coordinateTreasures
}

func removeCoordinateTreasureIndex(s []Coordinate, index int) []Coordinate {
	temp := make([]Coordinate, 0)
	temp = append(temp, s[:index]...)
	return append(temp, s[index+1:]...)
}

func isEqualTreasure(x int, y int) bool {
	return x == coordinateTreasure.X && y == coordinateTreasure.Y
}
