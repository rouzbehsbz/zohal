package archetype

import (
	"reflect"

	"github.com/rouzbehsbz/zohal/component"
)

type Column interface {
	Resize(length int)
	Len() int
	Remove(index int) bool
	Set(index int, value any) bool
	Get(index int) (any, bool)
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
