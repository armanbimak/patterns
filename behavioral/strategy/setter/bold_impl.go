package setter

import (
	"encoding/json"
	"fmt"
)

func NewAgreementBoldSetter() ISetter {
	return &agreementBoldSetter{}
}

type agreementBoldSetter struct{}

func (s *agreementBoldSetter) SetData(pattern []byte, data json.RawMessage) {
	fmt.Println("Setting bold data in given pattern.")
}
