package world

import "time"

type Stage = uint8

const (
	StartupStage Stage = iota
)

type Scheduler struct {
	systems map[Stage][]System
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		systems: make(map[Stage][]System),
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
	s.RunStage(world, StartupStage, 0)
}
