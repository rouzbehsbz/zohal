package component

import "reflect"

type ComponentRegistry struct {
	registry map[reflect.Type]int
	counter  int
}

func NewComponentRegistry() *ComponentRegistry {
	return &ComponentRegistry{
		registry: make(map[reflect.Type]int),
	}
}

func (c *ComponentRegistry) ComponentId(component any) int {
	componentType := reflect.TypeOf(component)
	componentId, ok := c.registry[componentType]
	if !ok {
		componentId = c.counter
		c.counter += 1
		c.registry[componentType] = componentId
	}

	return componentId
}
