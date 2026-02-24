package zurvan

type Entity struct {
	Index      int
	Generation int
}

func NewEntity(index, generation int) Entity {
	return Entity{
		Index:      index,
		Generation: generation,
	}
}
