package archetype

import (
	"reflect"

	"github.com/rouzbehsbz/zurvan/component"
)

type Column interface {
	Resize(length int)
	Len() int
	Remove(index int)
	Set(index int, value any)
	Get(index int) any
	AsSlice() any
}

type ColumnEntry struct {
	ComponentId int
	Column      Column
}

func NewColumnEntry(componentId int, elemType reflect.Type) ColumnEntry {
	return ColumnEntry{
		ComponentId: componentId,
		Column:      component.NewVector(elemType),
	}
}
