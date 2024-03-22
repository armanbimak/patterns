package agreement

import (
	"github.com/armanbimak/patterns/behavioral/strategy/generator"
	"github.com/armanbimak/patterns/behavioral/strategy/setter"
)

type IAgreement interface {
	Fill()
	Generate()

	SetFiller(setter setter.ISetter)
	SetGenerator(generator generator.IAgreementFileGenerator)
}
