package main

import "github.com/rouzbehsbz/zohal/world"

func main() {
	world := world.NewWorld()

	world.Spawn(
		Position{X: 2, Y: 3},
		Velocity{X: 1, Y: 1},
	)
}

type Position struct {
	X int
	Y int
}

type Velocity struct {
	X int
	Y int
}
