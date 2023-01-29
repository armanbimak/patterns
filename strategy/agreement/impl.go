package agreement

import (
	"encoding/json"
	"github.com/armanbimak/patterns/strategy/generator"
	"github.com/armanbimak/patterns/strategy/setter"
)

func NewAgreement(
	pattern []byte,
	data json.RawMessage,
	fileGenerator generator.IAgreementFileGenerator,
	setter setter.ISetter,
) IAgreement {

	return &agreement{
		pattern:    pattern,
		data:       data,
		generator:  fileGenerator,
		dataSetter: setter,
	}
}

type agreement struct {
	// Pattern docx document in bytes with predefined keys -> [{first-name}, {bank-account} ...],
	// which will be later filled with actual agreement data
	pattern []byte

	// Data agreement actual data, which will be set to predefined keys in docx pattern
	data json.RawMessage

	// Generator actual file generator docx or pdf
	generator generator.IAgreementFileGenerator // encapsulated algo family

	// DataSetter data filler in different formats (bold or cursive)
	dataSetter setter.ISetter // encapsulated algo family
}

func (a *agreement) SetFiller(setter setter.ISetter) {
	a.dataSetter = setter
}

func (a *agreement) SetGenerator(generator generator.IAgreementFileGenerator) {
	a.generator = generator
}

func (a *agreement) Fill() {
	a.dataSetter.SetData(a.pattern, a.data)
}

func (a *agreement) Generate() {
	a.generator.Generate()
}
