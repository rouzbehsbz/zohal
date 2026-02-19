package archetype

import "github.com/rouzbehsbz/zohal/entity"

type EntityLocation struct {
	Mask Mask
	Row  int
}

type ArchetypeAllocator struct {
	archetypes map[Mask]*Archetype
	locations  map[entity.Entity]EntityLocation
}

func NewArchetypeAllocator() *ArchetypeAllocator {
	return &ArchetypeAllocator{
		archetypes: make(map[Mask]*Archetype),
		locations:  make(map[entity.Entity]EntityLocation),
	}
}
