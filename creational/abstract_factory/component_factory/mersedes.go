package component_factory

import (
	component2 "github.com/armanbimak/patterns/creational/abstract_factory/component"
)

type MercedesFactory struct{}

func NewMercedesFactory() MercedesFactory {
	return MercedesFactory{}
}

func (m MercedesFactory) CreateEngine(types string) component2.IEngine {
	switch types {
	case "V8":
		return component2.V8EngineMercedes{}
	case "V6":
		return component2.V6EngineMercedes{}
	}

	return nil
}

func (m MercedesFactory) CreateWheel(types string) component2.IWheel {
	switch types {
	case "D24":
		return component2.D24WheelMercedes{}
	case "D26":
		return component2.D26WheelMercedes{}
	}

	return nil
}
