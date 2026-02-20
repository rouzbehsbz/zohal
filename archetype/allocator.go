package archetype

import (
	"reflect"
	"sort"

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

func (a *ArchetypeAllocator) AddComponents(entity entity.Entity, components ...any) {
	var mask Mask
	entries := make([]component.ComponentEntry, 0, len(components))

	for _, c := range components {
		id := a.registry.ComponentId(c)
		mask |= MaskBit(id)
		entries = append(entries, component.NewComponentEntry(id, reflect.TypeOf(c)))
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Id < entries[j].Id
	})

	targetArchetype, ok := a.archetypes[mask]
	if !ok {
		targetArchetype = NewArchetype(entries)
		a.archetypes[mask] = targetArchetype
	}

	location, exists := a.locations[entity]
	if !exists {
		a.AddNewEntity(entity, targetArchetype, mask, components)
		return
	}

	if mask == location.Mask {
		a.SetComponents(targetArchetype, location.Row, components)
		return
	}

	a.MoveEntity(entity, location, targetArchetype, mask)
}

func (a *ArchetypeAllocator) RemoveEntity(entity entity.Entity) {
	location := a.locations[entity]
	archetype := a.archetypes[location.Mask]

	archetype.RemoveEntity(location.Row)
	delete(a.locations, entity)
}

func (a *ArchetypeAllocator) AddNewEntity(entity entity.Entity, archetype *Archetype, mask Mask, components []any) {
	row := archetype.AddEntity(entity)

	a.SetComponents(archetype, row, components)

	a.locations[entity] = NewEntityLocation(mask, row)
}

func (a *ArchetypeAllocator) MoveEntity(entity entity.Entity, location EntityLocation, target *Archetype, newMask Mask) {
	source := a.archetypes[location.Mask]

	newRow := target.AddEntity(entity)

	source.MoveComponents(location.Row, newRow, target)
	source.RemoveEntity(location.Row)

	a.locations[entity] = NewEntityLocation(newMask, newRow)
}

func (a *ArchetypeAllocator) SetComponents(archetype *Archetype, row int, components []any) {
	for _, c := range components {
		id := a.registry.ComponentId(c)
		archetype.AddComponent(row, id, c)
	}
}

func (a *ArchetypeAllocator) MatchingArchetypes(componentIds ...int) []*Archetype {
	archetypes := []*Archetype{}

	for mask, archetype := range a.archetypes {
		queryMask := MaskBit(componentIds...)

		if MaskHasComponent(mask, queryMask) {
			archetypes = append(archetypes, archetype)
		}
	}

	return archetypes
}
