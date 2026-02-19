package archetype

import (
	"reflect"

	"github.com/rouzbehsbz/zohal/component"
	"github.com/rouzbehsbz/zohal/entity"
)

type ArchetypeAllocator struct {
	archetypes map[Mask]*Archetype
	locations  map[entity.Entity]EntityLocation

	registry *component.ComponentRegistry
}

func NewArchetypeAllocator(registry *component.ComponentRegistry) *ArchetypeAllocator {
	return &ArchetypeAllocator{
		archetypes: make(map[Mask]*Archetype),
		locations:  make(map[entity.Entity]EntityLocation),
		registry:   registry,
	}
}

func (a *ArchetypeAllocator) AddComponents(entity entity.Entity, components ...any) bool {
	var mask Mask
	entries := make([]component.ComponentEntry, 0, len(components))

	for _, c := range components {
		id := a.registry.ComponentId(c)
		mask |= MaskBit(id)
		entries = append(entries, component.NewComponentEntry(id, reflect.TypeOf(c)))
	}

	targetArchetype := a.Archetype(mask, entries)

	location, exists := a.locations[entity]
	if !exists {
		return a.AddNewEntity(entity, targetArchetype, mask, components)
	}

	if mask == location.Mask {
		return a.SetComponents(targetArchetype, location.Row, components)
	}

	return a.MoveEntity(entity, location, targetArchetype, mask)
}

func (a *ArchetypeAllocator) RemoveEntity(entity entity.Entity) bool {
	location, exists := a.locations[entity]
	if !exists {
		return false
	}

	archetype := a.archetypes[location.Mask]
	if archetype == nil {
		return false
	}

	archetype.RemoveEntity(location.Row)
	delete(a.locations, entity)

	return true
}

func (a *ArchetypeAllocator) AddNewEntity(entity entity.Entity, archetype *Archetype, mask Mask, components []any) bool {
	row := archetype.AddEntity(entity)

	if !a.SetComponents(archetype, row, components) {
		return false
	}

	a.locations[entity] = NewEntityLocation(mask, row)
	return true
}

func (a *ArchetypeAllocator) Archetype(mask Mask, entries []component.ComponentEntry) *Archetype {
	if arch, ok := a.archetypes[mask]; ok {
		return arch
	}

	arch := NewArchetype(entries)
	a.archetypes[mask] = arch

	return arch
}

func (a *ArchetypeAllocator) MoveEntity(entity entity.Entity, location EntityLocation, target *Archetype, newMask Mask) bool {
	source, ok := a.archetypes[location.Mask]
	if !ok {
		return false
	}

	newRow := target.AddEntity(entity)

	source.MoveComponents(location.Row, newRow, target)
	source.RemoveEntity(location.Row)

	a.locations[entity] = NewEntityLocation(newMask, newRow)

	return true
}

func (a *ArchetypeAllocator) SetComponents(archetype *Archetype, row int, components []any) bool {
	for _, c := range components {
		id := a.registry.ComponentId(c)
		if !archetype.AddComponent(row, id, c) {
			return false
		}
	}

	return true
}
