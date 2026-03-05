package zurvan

func QueryMany1[A any](world *World, fn func([]Entity, []A)) {
	componentId := DataIdFor[A](world.componentRegistry)

	archetypes := world.archetypeAllocator.MatchingArchetypes(componentId)
	for _, archetype := range archetypes {
		entities := archetype.Entities()
		column := archetype.Column(componentId)
		slice := column.AsSlice().([]A)

		fn(entities, slice)
	}
}

func QueryMany2[A, B any](world *World, fn func([]Entity, []A, []B)) {
	componentAId := DataIdFor[A](world.componentRegistry)
	componentBId := DataIdFor[B](world.componentRegistry)

	archetypes := world.archetypeAllocator.MatchingArchetypes(componentAId, componentBId)
	for _, archetype := range archetypes {
		entities := archetype.Entities()

		columnA := archetype.Column(componentAId)
		columnB := archetype.Column(componentBId)

		sliceA := columnA.AsSlice().([]A)
		sliceB := columnB.AsSlice().([]B)

		fn(entities, sliceA, sliceB)
	}
}

func QueryMany3[A, B, C any](world *World, fn func([]Entity, []A, []B, []C)) {
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

		fn(entities, sliceA, sliceB, sliceC)
	}
}

func QueryMany4[A, B, C, D any](world *World, fn func([]Entity, []A, []B, []C, []D)) {
	componentAId := DataIdFor[A](world.componentRegistry)
	componentBId := DataIdFor[B](world.componentRegistry)
	componentCId := DataIdFor[C](world.componentRegistry)
	componentDId := DataIdFor[D](world.componentRegistry)

	archetypes := world.archetypeAllocator.MatchingArchetypes(componentAId, componentBId, componentCId, componentDId)
	for _, archetype := range archetypes {
		entities := archetype.Entities()

		columnA := archetype.Column(componentAId)
		columnB := archetype.Column(componentBId)
		columnC := archetype.Column(componentCId)
		columnD := archetype.Column(componentDId)

		sliceA := columnA.AsSlice().([]A)
		sliceB := columnB.AsSlice().([]B)
		sliceC := columnC.AsSlice().([]C)
		sliceD := columnD.AsSlice().([]D)

		fn(entities, sliceA, sliceB, sliceC, sliceD)
	}
}

func QueryOne1[A any](world *World, entity Entity) *A {
	componentId := DataIdFor[A](world.componentRegistry)

	archetype, row := world.archetypeAllocator.MatchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil
	}

	column := archetype.Column(componentId)
	data := column.Get(row).(A)

	return &data
}

func QueryOne2[A, B any](world *World, entity Entity) (*A, *B) {
	componentAId := DataIdFor[A](world.componentRegistry)
	componentBId := DataIdFor[B](world.componentRegistry)

	archetype, row := world.archetypeAllocator.MatchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil
	}

	columnA := archetype.Column(componentAId)
	columnB := archetype.Column(componentBId)

	dataA := columnA.Get(row).(A)
	dataB := columnB.Get(row).(B)

	return &dataA, &dataB
}

func QueryOne3[A, B, C any](world *World, entity Entity) (*A, *B, *C) {
	componentAId := DataIdFor[A](world.componentRegistry)
	componentBId := DataIdFor[B](world.componentRegistry)
	componentCId := DataIdFor[C](world.componentRegistry)

	archetype, row := world.archetypeAllocator.MatchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil, nil
	}

	columnA := archetype.Column(componentAId)
	columnB := archetype.Column(componentBId)
	columnC := archetype.Column(componentCId)

	dataA := columnA.Get(row).(A)
	dataB := columnB.Get(row).(B)
	dataC := columnC.Get(row).(C)

	return &dataA, &dataB, &dataC
}

func QueryOne4[A, B, C, D any](world *World, entity Entity) (*A, *B, *C, *D) {
	componentAId := DataIdFor[A](world.componentRegistry)
	componentBId := DataIdFor[B](world.componentRegistry)
	componentCId := DataIdFor[C](world.componentRegistry)
	componentDId := DataIdFor[D](world.componentRegistry)

	archetype, row := world.archetypeAllocator.MatchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil, nil, nil
	}

	columnA := archetype.Column(componentAId)
	columnB := archetype.Column(componentBId)
	columnC := archetype.Column(componentCId)
	columnD := archetype.Column(componentDId)

	dataA := columnA.Get(row).(A)
	dataB := columnB.Get(row).(B)
	dataC := columnC.Get(row).(C)
	dataD := columnD.Get(row).(D)

	return &dataA, &dataB, &dataC, &dataD
}
