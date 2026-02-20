package main

import (
	"fmt"
	"time"

	"github.com/rouzbehsbz/zohal/world"
)

func main() {
	w := world.NewWorld(100 * time.Millisecond)

	w.AddSystems(
		&MovementSystem{},
		&PrintSystem{},
	)

	w.Spawn(
		Name("Rouzbeh"),
		Position{X: 0, Y: 0},
		Velocity{X: 1, Y: 2},
	)

	w.Spawn(
		Position{X: 0, Y: 0},
		Name("Shayan"),
		Velocity{X: 3, Y: -2},
	)

	w.Run()
}

type Name string

type Position struct {
	X, Y float64
}

type Velocity struct {
	X, Y float64
}

type MovementSystem struct{}

func (m *MovementSystem) Stage() world.Stage {
	return world.UpdateStage
}
func (m *MovementSystem) Update(w *world.World, dt time.Duration) {
	world.Query2[Position, Velocity](w, func(p *Position, v *Velocity) {
		p.X += v.X * dt.Seconds()
		p.Y += v.Y * dt.Seconds()
	})
}

type PrintSystem struct{}

func (p *PrintSystem) Stage() world.Stage {
	return world.UpdateStage
}
func (p *PrintSystem) Update(w *world.World, dt time.Duration) {
	world.Query2[Position, Name](w, func(p *Position, n *Name) {
		fmt.Printf("%s is at (%f, %f)\n", *n, p.X, p.Y)
	})
}
