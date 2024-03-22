package command

import (
	"github.com/armanbimak/patterns/behavioral/command/reciever"
)

type LightCommand struct {
	light reciever.ILight
}

func (l LightCommand) Execute() {
	l.light.On()
}

func (l LightCommand) Undo() {
	l.light.Off()
}

func NewLightCommand(light reciever.ILight) ICommand {
	return LightCommand{light: light}
}
