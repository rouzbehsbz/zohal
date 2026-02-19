package archetype

type Mask = uint32

func MaskBit(componentId int) Mask {
	return 1 << componentId
}
