package command

type ICommand interface {
	Execute()
	Undo()
}
