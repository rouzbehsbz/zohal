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
	for _, command := range c.commands {
		command.Execute(w)
	}

	c.commands = c.commands[:0]
}

type SpawnCommand struct {
	components []any
}

func NewSpawnCommand(components ...any) *SpawnCommand {
	return &SpawnCommand{
		components: components,
	}
}

func (s *SpawnCommand) Execute(w *World) {
	entity := w.entityAllocator.Create()
	w.archetypeAllocator.AddComponents(entity, s.components...)
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
