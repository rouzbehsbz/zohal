package world

import (
	"github.com/rouzbehsbz/zohal/component"
)

func Query1[A any](world *World, fn func(*A)) {
	componentId := component.ComponentIdFor[A](world.registry)

	archetypes := world.archetypeAllocator.MatchingArchetypes(componentId)
	for _, archetype := range archetypes {
		column := archetype.Column(componentId)
		slice := column.AsSlice().([]A)

		for i := range slice {
			fn(&slice[i])
		}
	}
}

func Query2[A, B any](world *World, fn func(*A, *B)) {
	componentAId := component.ComponentIdFor[A](world.registry)
	componentBId := component.ComponentIdFor[B](world.registry)

	archetypes := world.archetypeAllocator.MatchingArchetypes(componentAId, componentBId)
	for _, archetype := range archetypes {
		columnA := archetype.Column(componentAId)
		columnB := archetype.Column(componentBId)

		sliceA := columnA.AsSlice().([]A)
		sliceB := columnB.AsSlice().([]B)

		for i := range len(sliceA) {
			fn(&sliceA[i], &sliceB[i])
		}
	}
}
