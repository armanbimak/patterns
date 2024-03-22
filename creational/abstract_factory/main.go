package main

import (
	"github.com/armanbimak/patterns/creational/abstract_factory/component_factory"
)

func main() {
	bmw := component_factory.NewBMWFactory()
	mercedes := component_factory.NewMercedesFactory()

	bmwEngine := bmw.CreateEngine("V8")
	mercedesEngine := mercedes.CreateEngine("V6")

	bmwEngine.Power()
	mercedesEngine.Power()
}
