package zurvan

func Query1[A any](world *World, fn func(Entity, *A)) {
	componentId := DataIdFor[A](world.componentRegistry)

	archetypes := world.archetypeAllocator.MatchingArchetypes(componentId)
	for _, archetype := range archetypes {
		entities := archetype.Entities()
		column := archetype.Column(componentId)
		slice := column.AsSlice().([]A)

		for i := range len(entities) {
			fn(entities[i], &slice[i])
		}
	}
}

func Query2[A, B any](world *World, fn func(Entity, *A, *B)) {
	componentAId := DataIdFor[A](world.componentRegistry)
	componentBId := DataIdFor[B](world.componentRegistry)

	archetypes := world.archetypeAllocator.MatchingArchetypes(componentAId, componentBId)
	for _, archetype := range archetypes {
		entities := archetype.Entities()

		columnA := archetype.Column(componentAId)
		columnB := archetype.Column(componentBId)

		sliceA := columnA.AsSlice().([]A)
		sliceB := columnB.AsSlice().([]B)

		for i := range len(entities) {
			fn(entities[i], &sliceA[i], &sliceB[i])
		}
	}
}

func Query3[A, B, C any](world *World, fn func(Entity, *A, *B, *C)) {
	componentAId := DataIdFor[A](world.componentRegistry)
	componentBId := DataIdFor[B](world.componentRegistry)
	componentCId := DataIdFor[C](world.componentRegistry)

	archetypes := world.archetypeAllocator.MatchingArchetypes(componentAId, componentBId, componentCId)
	for _, archetype := range archetypes {
		entities := archetype.Entities()

		columnA := archetype.Column(componentAId)
		columnB := archetype.Column(componentBId)
		columnC := archetype.Column(componentCId)

		sliceA := columnA.AsSlice().([]A)
		sliceB := columnB.AsSlice().([]B)
		sliceC := columnC.AsSlice().([]C)

		for i := range len(entities) {
			fn(entities[i], &sliceA[i], &sliceB[i], &sliceC[i])
		}
	}
}
