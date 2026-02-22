package storage

import "reflect"

type Registry struct {
	registry map[reflect.Type]int
	counter  int
}

func NewRegistry() *Registry {
	return &Registry{
		registry: make(map[reflect.Type]int),
		counter:  0,
	}
}

func (r *Registry) dataId(dataType reflect.Type) int {
	dataId, ok := r.registry[dataType]
	if !ok {
		dataId = r.counter
		r.counter += 1
		r.registry[dataType] = dataId
	}

	return dataId
}

func (r *Registry) DataId(data any) int {
	dataType := reflect.TypeOf(data)

	return r.dataId(dataType)
}

func DataIdFor[T any](registry *Registry) int {
	dataType := reflect.TypeFor[T]()

	return registry.dataId(dataType)
}
