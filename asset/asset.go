package asset

type Coordinate struct {
	X int
	Y int
}

var (
	xGridLength = 8
	yGridLength = 6
)

//set initial coordinates
func init() {
	initObstacleCoordinates()
	setObstacleCoordinates()
	setTreasureCoordinates()
}
