package command

import "github.com/armanbimak/patterns/command/reciever"

type FanCommand struct {
	fan reciever.IFan
}

func NewFanCommand(fan reciever.IFan) *FanCommand {
	return &FanCommand{fan: fan}
}

func (f FanCommand) Execute() {
	f.fan.On()
	// will speed up slowly
	f.fan.SetSpeed(40)
}

func (f FanCommand) Undo() {
	f.fan.SetSpeed(0)
	f.fan.Off()
}
