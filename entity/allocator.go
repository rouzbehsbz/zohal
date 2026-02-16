package entity

type EntityAllocator struct {
	generations []uint32
	availables  []uint32
}

func NewEntityAllocator() *EntityAllocator {
	return &EntityAllocator{
		generations: []uint32{},
		availables:  []uint32{},
	}
}

func (e *EntityAllocator) Create() Entity {
	if len(e.availables) == 0 {
		index := len(e.generations)
		e.generations = append(e.generations, 0)

		return NewEntity(uint32(index), 0)
	}

	lastIndex := len(e.availables) - 1
	index := e.availables[lastIndex]

	e.availables = e.availables[:lastIndex]
	generation := e.generations[index]

	return NewEntity(uint32(index), generation)
}

func (e *EntityAllocator) Delete(entity Entity) bool {
	if e.IsAlive(entity) {
		e.availables = append(e.availables, entity.Index)
		e.generations[entity.Index] += 1

		return true
	}

	return false
}

func (e *EntityAllocator) IsAlive(entity Entity) bool {
	if int(entity.Index) >= len(e.generations) {
		return false
	}

	generation := e.generations[entity.Index]

	return generation == entity.Generation
}
