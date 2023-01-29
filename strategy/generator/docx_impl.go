package generator

import "fmt"

func NewAgreementDocxFileGenerator() IAgreementFileGenerator {
	return &agreementDocxFileGenerator{}
}

type agreementDocxFileGenerator struct{}

func (g *agreementDocxFileGenerator) Generate() {
	fmt.Println("Generating DOCX document.")
}
