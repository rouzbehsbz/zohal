package archetype

import (
	"github.com/rouzbehsbz/zurvan/component"
	"github.com/rouzbehsbz/zurvan/entity"
)

type Archetype struct {
	entities []entity.Entity
	columns  []ColumnEntry

	componentIndex map[int]int
}

func NewArchetype(entries []component.ComponentEntry) *Archetype {
	columns := []ColumnEntry{}
	componentIndex := make(map[int]int, len(entries))

	for _, entry := range entries {
		index := len(columns)

		columns = append(columns, NewColumnEntry(entry.Id, entry.ElemType))

		componentIndex[entry.Id] = index
	}

	return &Archetype{
		entities:       []entity.Entity{},
		columns:        columns,
		componentIndex: componentIndex,
	}
}

func (a *Archetype) IsEntityAlive(entity entity.Entity, row int) bool {
	if row >= len(a.entities) {
		return false
	}

	e := a.entities[row]

	return e.Index == entity.Index && e.Generation == entity.Generation
}

func (a *Archetype) AddEntity(entity entity.Entity) int {
	row := len(a.entities)
	a.entities = append(a.entities, entity)

	for _, entry := range a.columns {
		entry.Column.Resize(len(a.entities))
	}

	return row
}

func (a *Archetype) RemoveEntity(row int) (entity.Entity, int) {
	length := len(a.entities)
	if row >= length {
		return entity.Entity{}, -1
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

	return entity.Entity{}, -1
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

func (a *Archetype) Column(componentId int) Column {
	entry := a.columns[componentId]

	return entry.Column
}
