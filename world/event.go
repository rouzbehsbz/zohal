package world

import (
	"reflect"

	"github.com/rouzbehsbz/zurvan/storage"
)

type Events struct {
	columns map[int]storage.Column

	registry *storage.Registry
}

func NewEvents(registry *storage.Registry) *Events {
	return &Events{
		columns:  make(map[int]storage.Column),
		registry: registry,
	}
}

func (e *Events) Emit(event any) {
	eventId := e.registry.DataId(event)

	column, ok := e.columns[eventId]
	if !ok {
		column = storage.NewVector(reflect.TypeOf(event))
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
	eventId := storage.DataIdFor[T](w.events.registry)

	column, ok := w.events.columns[eventId]
	if !ok {
		return []T{}
	}

	slice := column.AsSlice()

	return slice.([]T)
}
