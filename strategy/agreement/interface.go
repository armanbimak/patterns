package agreement

import (
	"github.com/armanbimak/patterns/strategy/generator"
	"github.com/armanbimak/patterns/strategy/setter"
)

type IAgreement interface {
	Fill()
	Generate()

	SetFiller(setter setter.ISetter)
	SetGenerator(generator generator.IAgreementFileGenerator)
}
