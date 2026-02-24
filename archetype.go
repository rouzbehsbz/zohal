package zurvan

import (
	"reflect"
)

type ColumnEntry struct {
	ComponentId int
	Column      Column
}

func NewColumnEntry(componentId int, elemType reflect.Type) ColumnEntry {
	return ColumnEntry{
		ComponentId: componentId,
		Column:      NewVector(elemType),
	}
}

type Archetype struct {
	entities []Entity
	columns  []ColumnEntry

	componentIndex map[int]int
}

func NewArchetype(entries []ComponentEntry) *Archetype {
	columns := []ColumnEntry{}
	componentIndex := make(map[int]int, len(entries))

	for _, entry := range entries {
		index := len(columns)

		columns = append(columns, NewColumnEntry(entry.Id, entry.ElemType))

		componentIndex[entry.Id] = index
	}

	return &Archetype{
		entities:       []Entity{},
		columns:        columns,
		componentIndex: componentIndex,
	}
}

func (a *Archetype) IsEntityAlive(entity Entity, row int) bool {
	if row >= len(a.entities) {
		return false
	}

	e := a.entities[row]

	return e.Index == entity.Index && e.Generation == entity.Generation
}

func (a *Archetype) AddEntity(entity Entity) int {
	row := len(a.entities)
	a.entities = append(a.entities, entity)

	for _, entry := range a.columns {
		entry.Column.Resize(len(a.entities))
	}

	return row
}

func (a *Archetype) RemoveEntity(row int) (Entity, int) {
	length := len(a.entities)
	if row >= length {
		return Entity{}, -1
	}

	lastIndex := length - 1
	swapped := a.entities[lastIndex]
	a.entities[row] = swapped

	a.entities = a.entities[:lastIndex]

	for _, entry := range a.columns {
		entry.Column.Remove(row)
	}

	if row != lastIndex {
		return swapped, row
	}

	return Entity{}, -1
}

func (a *Archetype) AddComponent(row int, componentId int, component any) {
	columnIndex := a.componentIndex[componentId]
	entry := a.columns[columnIndex]
	entry.Column.Set(row, component)
}

func (a *Archetype) MoveComponents(row int, dstRow int, dstArchetype *Archetype) {
	for _, entry := range a.columns {
		component := entry.Column.Get(row)
		dstArchetype.AddComponent(dstRow, entry.ComponentId, component)
	}
}

func (a *Archetype) Entities() []Entity {
	return a.entities
}

func (a *Archetype) Column(componentId int) Column {
	entry := a.columns[componentId]

	return entry.Column
}
