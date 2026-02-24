package zurvan

import (
	"reflect"
	"sort"
)

type EntityLocation struct {
	Mask Mask
	Row  int
}

func NewEntityLocation(mask Mask, row int) EntityLocation {
	return EntityLocation{
		Mask: mask,
		Row:  row,
	}
}

type ComponentEntry struct {
	Id       int
	ElemType reflect.Type
}

func NewComponentEntry(id int, elemType reflect.Type) ComponentEntry {
	return ComponentEntry{
		Id:       id,
		ElemType: elemType,
	}
}

type ArchetypeAllocator struct {
	archetypes map[Mask]*Archetype
	locations  map[Entity]EntityLocation

	registry *Registry
}

func NewArchetypeAllocator(registry *Registry) *ArchetypeAllocator {
	return &ArchetypeAllocator{
		archetypes: make(map[Mask]*Archetype),
		locations:  make(map[Entity]EntityLocation),
		registry:   registry,
	}
}

func (a *ArchetypeAllocator) AddComponents(entity Entity, components ...any) {
	var mask Mask

	entries := make([]ComponentEntry, 0, len(components))

	for _, c := range components {
		id := a.registry.DataId(c)
		mask |= MaskBit(id)
		entries = append(entries, NewComponentEntry(id, reflect.TypeOf(c)))
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Id < entries[j].Id
	})

	location, exists := a.locations[entity]
	if exists {
		targetArchetype := a.archetypes[location.Mask]

		if MaskHasComponents(location.Mask, mask) {
			a.SetComponents(targetArchetype, location.Row, components)
			return
		}

		targetArchetype, ok := a.archetypes[mask]
		if !ok {
			targetArchetype = NewArchetype(entries)
			a.archetypes[mask] = targetArchetype
		}

		a.MoveEntity(entity, location, targetArchetype, mask)
		return
	}

	targetArchetype, ok := a.archetypes[mask]
	if !ok {
		targetArchetype = NewArchetype(entries)
		a.archetypes[mask] = targetArchetype
	}

	a.AddNewEntity(entity, targetArchetype, mask, components)
}

func (a *ArchetypeAllocator) RemoveEntity(entity Entity) {
	location := a.locations[entity]
	archetype := a.archetypes[location.Mask]

	archetype.RemoveEntity(location.Row)
	delete(a.locations, entity)
}

func (a *ArchetypeAllocator) AddNewEntity(entity Entity, archetype *Archetype, mask Mask, components []any) {
	row := archetype.AddEntity(entity)

	a.SetComponents(archetype, row, components)

	a.locations[entity] = NewEntityLocation(mask, row)
}

func (a *ArchetypeAllocator) MoveEntity(entity Entity, location EntityLocation, target *Archetype, newMask Mask) {
	source := a.archetypes[location.Mask]

	newRow := target.AddEntity(entity)

	source.MoveComponents(location.Row, newRow, target)
	source.RemoveEntity(location.Row)

	a.locations[entity] = NewEntityLocation(newMask, newRow)
}

func (a *ArchetypeAllocator) SetComponents(archetype *Archetype, row int, components []any) {
	for _, c := range components {
		id := a.registry.DataId(c)
		archetype.AddComponent(row, id, c)
	}
}

func (a *ArchetypeAllocator) MatchingArchetypes(componentIds ...int) []*Archetype {
	archetypes := []*Archetype{}

	for mask, archetype := range a.archetypes {
		queryMask := MaskBit(componentIds...)

		if MaskHasComponents(mask, queryMask) {
			archetypes = append(archetypes, archetype)
		}
	}

	return archetypes
}
