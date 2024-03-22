package initiator

import (
	command2 "github.com/armanbimak/patterns/behavioral/command/command"
	"sync"
)

type SmartVoiceController struct {
	commands    map[string]command2.ICommand
	lastCommand command2.ICommand

	mt *sync.Mutex // preferred simple rather than RW since lastCommand field synchronization
}

func (s *SmartVoiceController) Undo() {
	s.lastCommand.Undo()
}

func NewSmartVoiceController() *SmartVoiceController {
	return &SmartVoiceController{
		mt:          &sync.Mutex{},
		commands:    map[string]command2.ICommand{},
		lastCommand: command2.NewDummyCommand(),
	}
}

func (s *SmartVoiceController) OnVoiceCommandRecognized(command string) {
	s.mt.Lock()
	defer s.mt.Unlock()

	com := s.commands[command]
	s.lastCommand = com
	com.Execute()
}

func (s *SmartVoiceController) SetCommand(commandName string, commandObject command2.ICommand) {
	s.mt.Lock()
	defer s.mt.Unlock()

	s.commands[commandName] = commandObject
}
