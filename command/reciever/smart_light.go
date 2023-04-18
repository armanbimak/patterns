package reciever

import "fmt"

type ILight interface {
	On()
	Off()
}

type Light struct{}

func NewLight() ILight {
	return &Light{}
}

func (l Light) On() {
	fmt.Println("Light is ON")
}

func (l Light) Off() {
	fmt.Println("Light is OFF")
}
