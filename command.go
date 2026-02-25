package zurvan

type Command interface {
	Execute(w *World)
}

type Commands struct {
	commands []Command
}

func NewCommands() *Commands {
	return &Commands{
		commands: []Command{},
	}
}

func (c *Commands) AddCommand(command Command) {
	c.commands = append(c.commands, command)
}

func (c *Commands) Apply(w *World) {
	if len(c.commands) == 0 {
		return
	}

	for _, command := range c.commands {
		command.Execute(w)
	}

	c.commands = c.commands[:0]
}

type SetComponentsCommand struct {
	entity     Entity
	components []any
}

func NewSetComponentsCommand(entity Entity, components ...any) *SetComponentsCommand {
	return &SetComponentsCommand{
		entity:     entity,
		components: components,
	}
}

func (s *SetComponentsCommand) Execute(w *World) {
	w.archetypeAllocator.AddComponents(s.entity, s.components...)
}

type AddResourceCommand struct {
	resource any
}

func NewAddResourceCommand(resource any) *AddResourceCommand {
	return &AddResourceCommand{
		resource: resource,
	}
}

func (a *AddResourceCommand) Execute(w *World) {
	w.resources.AddResource(a.resource)
}

type DespawnCommand struct {
	entity Entity
}

func NewDespawnCommand(entity Entity) *DespawnCommand {
	return &DespawnCommand{
		entity: entity,
	}
}

func (d *DespawnCommand) Execute(w *World) {
	w.archetypeAllocator.RemoveEntity(d.entity)
	w.entityAllocator.Delete(d.entity)
}
