package zurvan

import "reflect"

type Resources struct {
	storage  map[int]any
	registry *Registry
}

func NewResources() *Resources {
	return &Resources{
		storage:  make(map[int]any),
		registry: NewRegistry(),
	}
}

func (r *Resources) AddResource(resource any) {
	resourceType := reflect.TypeOf(resource)
	resourceId := r.registry.dataId(resourceType)

	r.storage[resourceId] = resource
}

func Resource[T any](w *World) T {
	resourceId := DataIdFor[T](w.resources.registry)

	resource, ok := w.resources.storage[resourceId]
	if !ok {
		var defaultVal T

		resource = defaultVal
		w.resources.storage[resourceId] = resource
	}

	return resource.(T)
}
