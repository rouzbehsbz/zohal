package zurvan

import (
	"time"
)

type World struct {
	entityAllocator    *EntityAllocator
	archetypeAllocator *ArchetypeAllocator

	scheduler *Scheduler
	commands  *Commands
	events    *Events
	resources *Resources

	componentRegistry *Registry
}

func NewWorld(tickRate time.Duration) *World {
	componentRegistry := NewRegistry()
	eventRegistry := NewRegistry()

	commands := NewCommands()
	events := NewEvents(eventRegistry)

	return &World{
		entityAllocator:    NewEntityAllocator(),
		archetypeAllocator: NewArchetypeAllocator(componentRegistry),
		scheduler:          NewScheduler(commands, events, tickRate),
		commands:           commands,
		events:             events,
		resources:          NewResources(),
		componentRegistry:  componentRegistry,
	}
}

func (w *World) AddSystems(stage Stage, systems []System) {
	for _, system := range systems {
		w.scheduler.AddSystem(stage, system)
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
