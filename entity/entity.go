package entity

type Entity struct {
	Index      uint32
	Generation uint32
}

func NewEntity(index, generation uint32) Entity {
	return Entity{
		Index:      index,
		Generation: generation,
	}
}
