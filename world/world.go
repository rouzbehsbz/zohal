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
	commands           *Commands
	registry           *component.ComponentRegistry
}

func NewWorld(tickRate time.Duration) *World {
	commands := NewCommands()
	registry := component.NewComponentRegistry()

	return &World{
		entityAllocator:    entity.NewEntityAllocator(),
		archetypeAllocator: archetype.NewArchetypeAllocator(registry),
		scheduler:          NewScheduler(commands, tickRate),
		commands:           commands,
		registry:           registry,
	}
}

func (w *World) AddSystems(systems ...System) {
	for _, system := range systems {
		w.scheduler.AddSystem(system)
	}
}

func (w *World) AddCommands(commands ...Command) {
	for _, command := range commands {
		w.commands.AddCommand(command)
	}
}

func (w *World) Run() {
	w.scheduler.Run(w)
}
