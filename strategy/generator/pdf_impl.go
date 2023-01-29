package generator

import "fmt"

func NewAgreementPDFFileGenerator() IAgreementFileGenerator {
	return &agreementPDFFileGenerator{}
}

type agreementPDFFileGenerator struct{}

func (g *agreementPDFFileGenerator) Generate() {
	fmt.Println("Generating PDF document.")
}
