package zurvan

import (
	"reflect"
)

type Events struct {
	columns map[int]Column

	registry *Registry
}

func NewEvents(registry *Registry) *Events {
	return &Events{
		columns:  make(map[int]Column),
		registry: registry,
	}
}

func (e *Events) Emit(event any) {
	eventId := e.registry.DataId(event)

	column, ok := e.columns[eventId]
	if !ok {
		column = NewVector(reflect.TypeOf(event))
		e.columns[eventId] = column
	}

	column.Push(event)
}

func (e *Events) Clear() {
	for _, column := range e.columns {
		column.Resize(0)
	}
}

func OnEvent[T any](w *World) []T {
	eventId := DataIdFor[T](w.events.registry)

	column, ok := w.events.columns[eventId]
	if !ok {
		return []T{}
	}

	slice := column.AsSlice()

	return slice.([]T)
}
