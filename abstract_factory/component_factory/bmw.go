package component_factory

import "github.com/armanbimak/patterns/abstract_factory/component"

type BMWFactory struct{}

func NewBMWFactory() BMWFactory {
	return BMWFactory{}
}

func (B BMWFactory) CreateEngine(types string) component.IEngine {
	switch types {
	case "V8":
		return component.V8EngineBMW{}
	case "V6":
		return component.V6EngineBMW{}
	}

	return nil
}

func (B BMWFactory) CreateWheel(types string) component.IWheel {
	switch types {
	case "D24":
		return component.D24WheelBMW{}
	case "D26":
		return component.D26WheelBMW{}
	}

	return nil
}
