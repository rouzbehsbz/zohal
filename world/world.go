package world

import (
	"time"

	"github.com/rouzbehsbz/zurvan/archetype"
	"github.com/rouzbehsbz/zurvan/entity"
	"github.com/rouzbehsbz/zurvan/storage"
)

type World struct {
	entityAllocator    *entity.EntityAllocator
	archetypeAllocator *archetype.ArchetypeAllocator

	scheduler *Scheduler
	commands  *Commands
	events    *Events

	componentRegistry *storage.Registry
}

func NewWorld(tickRate time.Duration) *World {
	componentRegistry := storage.NewRegistry()
	eventRegistry := storage.NewRegistry()

	commands := NewCommands()
	events := NewEvents(eventRegistry)

	return &World{
		entityAllocator:    entity.NewEntityAllocator(),
		archetypeAllocator: archetype.NewArchetypeAllocator(componentRegistry),
		scheduler:          NewScheduler(commands, events, tickRate),
		commands:           commands,
		events:             events,
		componentRegistry:  componentRegistry,
	}
}

func (w *World) AddSystems(systems ...System) {
	for _, system := range systems {
		w.scheduler.AddSystem(system)
	}
}

func (w *World) PushCommands(commands ...Command) {
	for _, command := range commands {
		w.commands.AddCommand(command)
	}
}

func (w *World) EmitEvents(events ...any) {
	for _, event := range events {
		w.events.Emit(event)
	}
}

func (w *World) Run() {
	w.scheduler.Run(w)
}
