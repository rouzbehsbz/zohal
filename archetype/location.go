package archetype

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
