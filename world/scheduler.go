package world

import (
	"time"
)

type Stage = uint8

const (
	StartupStage Stage = iota
	PreUpdateStage
	FixedUpdateStage
	UpdateStage
)

type Scheduler struct {
	systems map[Stage][]System

	tickRate    time.Duration
	accumulator time.Duration
}

func NewScheduler(tickRate time.Duration) *Scheduler {
	return &Scheduler{
		systems:     make(map[Stage][]System),
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

func (s *Scheduler) AddSystem(system System) {
	stage := system.Stage()
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

		sleepTime := s.tickRate - time.Since(now)
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		}
	}
}
