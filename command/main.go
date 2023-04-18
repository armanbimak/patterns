package main

import (
	"github.com/armanbimak/patterns/command/command"
	"github.com/armanbimak/patterns/command/initiator"
	"github.com/armanbimak/patterns/command/reciever"
)

func main() {

	fan := reciever.NewFan()
	thermostat := reciever.NewThermostat()
	light := reciever.NewLight()

	fanCommand := command.NewFanCommand(fan)
	lightCommand := command.NewLightCommand(light)
	thermostatCommand := command.NewThermostatCommand(thermostat)
	macroCommand := command.NewMacroCommand(lightCommand, thermostatCommand)

	smartVoiceController := initiator.NewSmartVoiceController()

	smartVoiceController.SetCommand("fan", fanCommand)
	smartVoiceController.SetCommand("thermostat", thermostatCommand)
	smartVoiceController.SetCommand("light", lightCommand)
	smartVoiceController.SetCommand("macro", macroCommand)

	smartVoiceController.OnVoiceCommandRecognized("fan")
	smartVoiceController.OnVoiceCommandRecognized("thermostat")
	smartVoiceController.OnVoiceCommandRecognized("macro")

	smartVoiceController.Undo()
}
