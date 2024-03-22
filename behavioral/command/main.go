package main

import (
	command2 "github.com/armanbimak/patterns/behavioral/command/command"
	"github.com/armanbimak/patterns/behavioral/command/initiator"
	reciever2 "github.com/armanbimak/patterns/behavioral/command/reciever"
)

func main() {

	fan := reciever2.NewFan()
	thermostat := reciever2.NewThermostat()
	light := reciever2.NewLight()

	fanCommand := command2.NewFanCommand(fan)
	lightCommand := command2.NewLightCommand(light)
	thermostatCommand := command2.NewThermostatCommand(thermostat)
	macroCommand := command2.NewMacroCommand(lightCommand, thermostatCommand)

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
