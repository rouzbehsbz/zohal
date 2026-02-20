package world

import (
	"time"

	"github.com/rouzbehsbz/zohal/archetype"
	"github.com/rouzbehsbz/zohal/component"
	"github.com/rouzbehsbz/zohal/entity"
)

type World struct {
	entityAllocator    *entity.EntityAllocator
	archetypeAllocator *archetype.ArchetypeAllocator
	scheduler          *Scheduler
	registry           *component.ComponentRegistry
}

func NewWorld(tickRate time.Duration) *World {
	registry := component.NewComponentRegistry()

	return &World{
		entityAllocator:    entity.NewEntityAllocator(),
		archetypeAllocator: archetype.NewArchetypeAllocator(registry),
		scheduler:          NewScheduler(tickRate),
		registry:           registry,
	}
}

func (w *World) AddSystems(systems ...System) {
	for _, system := range systems {
		w.scheduler.AddSystem(system)
	}
}

func (w *World) Spawn(components ...any) {
	e := w.entityAllocator.Create()
	w.archetypeAllocator.AddComponents(e, components...)
}

func (w *World) Run() {
	w.scheduler.Run(w)
}
