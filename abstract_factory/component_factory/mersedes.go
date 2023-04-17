package component_factory

import "github.com/armanbimak/patterns/abstract_factory/component"

type MercedesFactory struct{}

func NewMercedesFactory() MercedesFactory {
	return MercedesFactory{}
}

func (m MercedesFactory) CreateEngine(types string) component.IEngine {
	switch types {
	case "V8":
		return component.V8EngineMercedes{}
	case "V6":
		return component.V6EngineMercedes{}
	}

	return nil
}

func (m MercedesFactory) CreateWheel(types string) component.IWheel {
	switch types {
	case "D24":
		return component.D24WheelMercedes{}
	case "D26":
		return component.D26WheelMercedes{}
	}

	return nil
}
