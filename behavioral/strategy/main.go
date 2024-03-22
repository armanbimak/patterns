package main

import (
	"encoding/json"
	"fmt"
	"github.com/armanbimak/patterns/behavioral/strategy/agreement"
	generator2 "github.com/armanbimak/patterns/behavioral/strategy/generator"
	setter2 "github.com/armanbimak/patterns/behavioral/strategy/setter"
)

func main() {
	docxGen := generator2.NewAgreementDocxFileGenerator()
	pdfGen := generator2.NewAgreementPDFFileGenerator()

	boldFiller := setter2.NewAgreementBoldSetter()
	cursiveFiller := setter2.NewAgreementCursiveSetter()

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
