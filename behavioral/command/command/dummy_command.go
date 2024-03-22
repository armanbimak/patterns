package command

type DummyCommand struct{}

func NewDummyCommand() ICommand {
	return &DummyCommand{}
}

func (d DummyCommand) Execute() {}

func (d DummyCommand) Undo() {}
