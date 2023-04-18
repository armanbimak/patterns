package reciever

import "fmt"

type IFan interface {
	On()
	Off()

	SetSpeed(speed int)
}

type Fan struct{}

func NewFan() IFan {
	return &Fan{}
}

func (f Fan) On() {
	fmt.Println("Fan is ON")
}

func (f Fan) Off() {
	fmt.Println("Fan is OFF")
}

func (f Fan) SetSpeed(speed int) {
	fmt.Println(fmt.Sprintf("Fan speed set to %d per second", speed))
}
