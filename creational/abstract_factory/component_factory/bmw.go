package component_factory

import (
	component2 "github.com/armanbimak/patterns/creational/abstract_factory/component"
)

type BMWFactory struct{}

func NewBMWFactory() BMWFactory {
	return BMWFactory{}
}

func (B BMWFactory) CreateEngine(types string) component2.IEngine {
	switch types {
	case "V8":
		return component2.V8EngineBMW{}
	case "V6":
		return component2.V6EngineBMW{}
	}

	return nil
}

func (B BMWFactory) CreateWheel(types string) component2.IWheel {
	switch types {
	case "D24":
		return component2.D24WheelBMW{}
	case "D26":
		return component2.D26WheelBMW{}
	}

	return nil
}
