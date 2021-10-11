package core

import (
	"fmt"
	"treasure-hunt/asset"
)

func Run() {
	//loop action based on total treasure possibilies
	for {
		if len(asset.GetCoordinateTreasures()) == 0 {
			break
		}
		// core action, take this out if you want to run without loop
		initAction()
	}

	fmt.Println("\n\nThanks for Watching")
	fmt.Println()
}
