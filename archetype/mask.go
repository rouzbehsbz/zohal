package archetype

type Mask = uint64

func MaskBit(componentIds ...int) Mask {
	var mask Mask

	for _, componentId := range componentIds {
		mask |= 1 << componentId
	}

	return mask
}

func MaskHasComponent(mask, query Mask) bool {
	return mask&query == query
}
