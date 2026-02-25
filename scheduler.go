package zurvan

import (
	"time"
)

type Stage = uint8

const (
	StartupStage Stage = iota
	PreUpdateStage
	FixedUpdateStage
	UpdateStage
	PostUpdateStage
)

type System interface {
	Update(w *World, dt time.Duration)
}

func BuildStageSystems(stage Stage, systems ...System) (Stage, []System) {
	return stage, systems
}

type Scheduler struct {
	systems  map[Stage][]System
	commands *Commands
	events   *Events

	tickRate    time.Duration
	accumulator time.Duration
}

func NewScheduler(commands *Commands, events *Events, tickRate time.Duration) *Scheduler {
	return &Scheduler{
		systems:     make(map[Stage][]System),
		commands:    commands,
		events:      events,
		tickRate:    tickRate,
		accumulator: 0,
	}
}

func (s *Scheduler) Stage(stage Stage) []System {
	systems, ok := s.systems[stage]
	if !ok {
		systems = []System{}
		s.systems[stage] = systems
	}

	return systems
}

func (s *Scheduler) AddSystem(stage Stage, system System) {
	systems := s.Stage(stage)
	systems = append(systems, system)
	s.systems[stage] = systems
}

func (s *Scheduler) RunStage(world *World, stage Stage, dt time.Duration) {
	systems := s.Stage(stage)

	for _, system := range systems {
		system.Update(world, dt)
	}
}

func (s *Scheduler) Run(world *World) {
	last := time.Now()

	s.RunStage(world, StartupStage, 0)

	for {
		now := time.Now()
		frameTime := now.Sub(last)
		last = now

		s.accumulator += frameTime

		s.RunStage(world, PreUpdateStage, frameTime)

		for s.accumulator >= s.tickRate {
			s.RunStage(world, FixedUpdateStage, s.tickRate)
			s.accumulator -= s.tickRate
		}

		s.RunStage(world, UpdateStage, frameTime)
		s.RunStage(world, PostUpdateStage, frameTime)

		s.commands.Apply(world)
		s.events.Clear()

		sleepTime := s.tickRate - time.Since(now)
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		}
	}
}
