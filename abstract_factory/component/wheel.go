package component

import "fmt"

type D24WheelBMW struct{}

func (d D24WheelBMW) Spin() {
	fmt.Println("D24 BMW Wheel Spin")
}

type D26WheelBMW struct{}

func (d D26WheelBMW) Spin() {
	fmt.Println("D26 BMW Wheel Spin")
}

type D24WheelMercedes struct{}

func (d D24WheelMercedes) Spin() {
	fmt.Println("D24 Mercedes Wheel Spin")
}

type D26WheelMercedes struct{}

func (d D26WheelMercedes) Spin() {
	fmt.Println("D26 Mercedes Wheel Spin")
}
