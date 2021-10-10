package asset

type Coordinate struct {
	X int
	Y int
}

var (
	xGridLength = 8
	yGridLength = 6
)

func init() {
	initObstacleCoordinates()
	setObstacleCoordinates()
	setTreasureCoordinates()
}
