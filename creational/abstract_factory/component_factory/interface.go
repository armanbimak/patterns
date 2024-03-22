package component_factory

import (
	"github.com/armanbimak/patterns/creational/abstract_factory/component"
)

type IVehicleComponentFactory interface {
	CreateEngine(types string) component.IEngine
	CreateWheel(types string) component.IWheel
}
