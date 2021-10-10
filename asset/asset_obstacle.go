package asset

var coordinateObstacles = map[Coordinate]bool{}

func initObstacleCoordinates() {
	coordinateObstacles = map[Coordinate]bool{
		{2, 2}: true,
		{3, 2}: true,
		{4, 2}: true,
		{4, 3}: true,
		{6, 3}: true,
		{2, 4}: true,
	}
}

func setObstacleCoordinates() {
	for i := 0; i < yGridLength; i++ {
		for j := 0; j < xGridLength; j++ {
			if i == 0 || i == yGridLength-1 || j == 0 || j == xGridLength-1 {
				coordinateObstacles[Coordinate{j, i}] = true
			}
		}
	}
}

func GetCoordinateObstacle() map[Coordinate]bool {
	return coordinateObstacles
}
