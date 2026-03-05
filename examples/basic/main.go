package main

import (
	"fmt"
	"time"

	"github.com/rouzbehsbz/zurvan"
)

func main() {
	world := zurvan.NewWorld(16 * time.Millisecond)

	world.AddSystems(
		zurvan.BuildStageSystems(zurvan.UpdateStage,
			&MovementSystem{},
			&RespawnSystem{},
		),
	)

	e := world.Spawn()

	world.PushCommands(
		zurvan.NewSetComponentsCommand(
			e,
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
	sec := dt.Seconds()

	zurvan.QueryMany2[Position, Velocity](w, func(entities []zurvan.Entity, p []Position, v []Velocity) {
		for i, e := range entities {
			p[i].X += v[i].X * sec
			p[i].Y += v[i].Y * sec

			fmt.Printf("pos: %v\n", p[i])

			if p[i].X > 50 && p[i].Y > 50 {
				w.EmitEvents(
					DeathEvent{
						Entity: e,
					},
				)
			}
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
