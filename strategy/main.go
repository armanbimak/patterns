package main

import (
	"encoding/json"
	"fmt"
	"github.com/armanbimak/patterns/strategy/agreement"
	"github.com/armanbimak/patterns/strategy/generator"
	"github.com/armanbimak/patterns/strategy/setter"
)

func main() {
	docxGen := generator.NewAgreementDocxFileGenerator()
	pdfGen := generator.NewAgreementPDFFileGenerator()

	boldFiller := setter.NewAgreementBoldSetter()
	cursiveFiller := setter.NewAgreementCursiveSetter()

	pattern := make([]byte, 0)
	data := make(json.RawMessage, 0)

	agr := agreement.NewAgreement(pattern, data, docxGen, boldFiller)

	agr.Fill()
	agr.Generate()

	fmt.Println("changing algo families in RUNTIME")
	agr.SetGenerator(pdfGen)
	agr.SetFiller(cursiveFiller)

	agr.Fill()
	agr.Generate()
}
