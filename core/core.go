package core

import (
	"fmt"
	"treasure-hunt/asset"
)

func Run() {
	for {
		if len(asset.GetCoordinateTreasures()) == 0 {
			break
		}
		initAction()
	}

	fmt.Println("\n\nThanks for Playing")
	fmt.Println()
}
