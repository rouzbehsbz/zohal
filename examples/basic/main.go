package main

import (
	"fmt"
	"time"

	"github.com/rouzbehsbz/zurvan"
)

func main() {
	world := zurvan.NewWorld(16 * time.Millisecond)

	world.AddSystems(
		&MovementSystem{},
		&RespawnSystem{},
	)

	world.PushCommands(
		zurvan.NewSpawnCommand(
			Position{X: 0, Y: 0},
			Velocity{X: 10, Y: 10},
		),
	)

	world.Run()
}

type Position struct {
	X, Y float64
}

type Velocity struct {
	X, Y float64
}

type DeathEvent struct {
	Entity zurvan.Entity
}

type MovementSystem struct{}

func (m *MovementSystem) Stage() zurvan.Stage {
	return zurvan.UpdateStage
}
func (m *MovementSystem) Update(w *zurvan.World, dt time.Duration) {
	zurvan.Query2[Position, Velocity](w, func(e zurvan.Entity, p *Position, v *Velocity) {
		dt := dt.Seconds()

		p.X += v.X * dt
		p.Y += v.Y * dt

		fmt.Printf("Position (%f, %f)\n", p.X, p.Y)

		if p.X > 50 && p.Y > 50 {
			fmt.Printf("Entering death zone\n")

			w.EmitEvents(
				DeathEvent{Entity: e},
			)
		}
	})
}

type RespawnSystem struct{}

func (m *RespawnSystem) Stage() zurvan.Stage {
	return zurvan.UpdateStage
}
func (m *RespawnSystem) Update(w *zurvan.World, dt time.Duration) {
	events := zurvan.OnEvent[DeathEvent](w)

	for _, event := range events {
		fmt.Printf("Respawning ...\n")

		w.PushCommands(
			zurvan.NewSetComponentsCommand(event.Entity,
				Position{X: 0, Y: 0},
			),
		)
	}
}
