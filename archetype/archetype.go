package archetype

import (
	"github.com/rouzbehsbz/zohal/component"
	"github.com/rouzbehsbz/zohal/entity"
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

func (a *Archetype) RemoveEntity(row int) (entity.Entity, int, bool) {
	length := len(a.entities)
	if row >= length {
		return entity.Entity{}, 0, false
	}

	lastIndex := length - 1
	swapped := a.entities[lastIndex]
	a.entities[row] = swapped

	a.entities = a.entities[:lastIndex]

	for _, entry := range a.columns {
		if ok := entry.Column.Remove(row); !ok {
			return entity.Entity{}, 0, false
		}
	}

	if row != lastIndex {
		return swapped, row, true
	}

	return entity.Entity{}, 0, true
}

func (a *Archetype) AddComponent(row int, componentId int, component any) bool {
	columnIndex, ok := a.componentIndex[componentId]
	if !ok {
		return false
	}

	entry := a.columns[columnIndex]

	return entry.Column.Set(row, component)
}

func (a *Archetype) MoveComponents(row int, dstRow int, dstArchetype *Archetype) bool {
	for _, entry := range a.columns {
		component, ok := entry.Column.Get(row)
		if !ok {
			return false
		}

		if ok := dstArchetype.AddComponent(dstRow, entry.ComponentId, component); !ok {
			return false
		}
	}

	return true
}
