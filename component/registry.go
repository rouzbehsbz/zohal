package component

import "reflect"

type ComponentRegistry struct {
	registry map[reflect.Type]int
	counter  int
}

func NewComponentRegistry() *ComponentRegistry {
	return &ComponentRegistry{
		registry: make(map[reflect.Type]int),
		counter:  0,
	}
}

func (c *ComponentRegistry) componentId(componentType reflect.Type) int {
	componentId, ok := c.registry[componentType]
	if !ok {
		componentId = c.counter
		c.counter += 1
		c.registry[componentType] = componentId
	}

	return componentId
}

func (c *ComponentRegistry) ComponentId(component any) int {
	componentType := reflect.TypeOf(component)

	return c.componentId(componentType)
}

func ComponentIdFor[T any](registry *ComponentRegistry) int {
	componentType := reflect.TypeFor[T]()

	return registry.componentId(componentType)
}
