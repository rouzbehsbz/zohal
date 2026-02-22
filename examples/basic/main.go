package main

import (
	"fmt"
	"time"

	"github.com/rouzbehsbz/zurvan/entity"
	"github.com/rouzbehsbz/zurvan/world"
)

func main() {
	w := world.NewWorld(100 * time.Millisecond)

	w.AddSystems(
		&TakeDamageSystem{},
		&RespawnSystem{},
	)

	w.PushCommands(
		world.NewSpawnCommand(
			Health{Amount: 100, Max: 100},
		),
	)

	w.Run()
}

type Health struct {
	Amount int
	Max    int
}

type DeathEvent struct {
	Entity entity.Entity
}

type TakeDamageSystem struct{}

func (t *TakeDamageSystem) Stage() world.Stage {
	return world.UpdateStage
}
func (t *TakeDamageSystem) Update(w *world.World, dt time.Duration) {
	world.Query1[Health](w, func(e entity.Entity, h *Health) {
		h.Amount -= 10

		fmt.Printf("entity %d health: %d\n", e.Index, h.Amount)

		if h.Amount <= 0 {
			w.EmitEvents(
				DeathEvent{Entity: e},
			)
		}
	})
}

type RespawnSystem struct{}

func (r *RespawnSystem) Stage() world.Stage {
	return world.UpdateStage
}
func (r *RespawnSystem) Update(w *world.World, dt time.Duration) {
	events := world.OnEvent[DeathEvent](w)

	for _, e := range events {
		fmt.Printf("entity %d has died. respawning ...\n", e.Entity.Index)

		w.PushCommands(
			world.NewSetComponentsCommand(e.Entity,
				Health{Amount: 100, Max: 100},
			),
		)
	}
}
