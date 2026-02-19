package component

import "reflect"

type ComponentEntry struct {
	Id       int
	ElemType reflect.Type
}

func NewComponentEntry(id int, elemType reflect.Type) ComponentEntry {
	return ComponentEntry{
		Id:       id,
		ElemType: elemType,
	}
}
