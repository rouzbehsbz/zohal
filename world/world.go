package world

import (
	"github.com/rouzbehsbz/zohal/archetype"
	"github.com/rouzbehsbz/zohal/component"
	"github.com/rouzbehsbz/zohal/entity"
)

type World struct {
	entityAllocator    *entity.EntityAllocator
	archetypeAllocator *archetype.ArchetypeAllocator
}

func NewWorld() *World {
	registry := component.NewComponentRegistry()

	return &World{
		entityAllocator:    entity.NewEntityAllocator(),
		archetypeAllocator: archetype.NewArchetypeAllocator(registry),
	}
}

func (w *World) Spawn(components ...any) {
	e := w.entityAllocator.Create()
	w.archetypeAllocator.AddComponents(e, components)
}

func (w *World) Despawn(entity entity.Entity) {
	w.entityAllocator.Delete(entity)
	w.archetypeAllocator.RemoveEntity(entity)
}
