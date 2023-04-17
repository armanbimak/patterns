package main

import (
	componentFactory "github.com/armanbimak/patterns/abstract_factory/component_factory"
)

func main() {
	bmw := componentFactory.NewBMWFactory()
	mercedes := componentFactory.NewMercedesFactory()

	bmwEngine := bmw.CreateEngine("V8")
	mercedesEngine := mercedes.CreateEngine("V6")

	bmwEngine.Power()
	mercedesEngine.Power()
}
