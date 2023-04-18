package command

type MacroCommand struct {
	commands []ICommand
}

func NewMacroCommand(commands ...ICommand) ICommand {
	return &MacroCommand{commands: commands}
}

func (m MacroCommand) Execute() {
	for _, command := range m.commands {
		command.Execute()
	}
}

func (m MacroCommand) Undo() {
	// reversed
	for i := len(m.commands) - 1; i >= 0; i-- {
		m.commands[i].Undo()
	}
}
