package initiator

import "github.com/armanbimak/patterns/command/command"

type ISmartVoiceController interface {
	OnVoiceCommandRecognized(command string)
	SetCommand(commandName string, commandObject command.ICommand)
	Undo()
}
