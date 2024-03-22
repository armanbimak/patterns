package reciever

import "fmt"

type IThermostat interface {
	SetTemperature(cells int)
}

type Thermostat struct{}

func NewThermostat() IThermostat {
	return &Thermostat{}
}

func (t Thermostat) SetTemperature(cells int) {
	fmt.Println(fmt.Sprintf("IThermostat - temperature %d C in set", cells))
}
