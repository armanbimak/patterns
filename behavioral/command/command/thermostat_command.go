package command

import (
	"github.com/armanbimak/patterns/behavioral/command/reciever"
)

type ThermostatCommand struct {
	thermostat reciever.IThermostat
}

func NewThermostatCommand(thermostat reciever.IThermostat) ICommand {
	return &ThermostatCommand{thermostat: thermostat}
}

func (t ThermostatCommand) Execute() {
	t.thermostat.SetTemperature(60)
}

func (t ThermostatCommand) Undo() {
	t.thermostat.SetTemperature(0)
}
