package zurvan

type EntityAllocator struct {
	generations []int
	availables  []int
}

func NewEntityAllocator() *EntityAllocator {
	return &EntityAllocator{
		generations: []int{},
		availables:  []int{},
	}
}

func (e *EntityAllocator) Create() Entity {
	if len(e.availables) == 0 {
		index := len(e.generations)
		e.generations = append(e.generations, 0)

		return NewEntity(index, 0)
	}

	lastIndex := len(e.availables) - 1
	index := e.availables[lastIndex]

	e.availables = e.availables[:lastIndex]
	generation := e.generations[index]

	return NewEntity(index, generation)
}

func (e *EntityAllocator) Delete(entity Entity) {
	if e.isAlive(entity) {
		e.availables = append(e.availables, entity.Index)
		e.generations[entity.Index] += 1
	}
}

func (e *EntityAllocator) isAlive(entity Entity) bool {
	if entity.Index >= len(e.generations) {
		return false
	}

	generation := e.generations[entity.Index]

	return generation == entity.Generation
}
